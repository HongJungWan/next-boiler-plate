package router

import (
	"context"
	"eCommerce/service/mysql"
	. "eCommerce/types"
	. "eCommerce/types/err"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type MySQLRouter struct {
	router       *Router
	mysqlService *mysql.MySQLService
}

func NewMySQLRouter(router *Router, mysqlService *mysql.MySQLService) {
	m := &MySQLRouter{
		router:       router,
		mysqlService: mysqlService,
	}

	baseUri := "/mysql"

	m.router.GET(baseUri+"/health", m.health)

	m.router.GET(baseUri+"/user-bucket", m.userBucket)                // 장바구니에 대한 정보를 가져오기
	m.router.GET(baseUri+"/content", m.content)                       // 상품 정보를 가져오는 정보
	m.router.GET(baseUri+"/user-bucket-history", m.userBucketHistory) // 유저의 구매 이력 정보

	m.router.POST(baseUri+"/create-user", m.createUser)       // 유저 데이터 생성
	m.router.POST(baseUri+"/create-content", m.createContent) // content 데이터 생성
	m.router.POST(baseUri+"/buy", m.buy)                      // history 데이터 생성
	m.router.POST(baseUri+"/bucket", m.bucket)                // 장바구니 데이터 생성
}

func (m *MySQLRouter) bucket(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mysqlService.PostBucketRequest(req.User, req.Content); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}

}

func (m *MySQLRouter) createUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mysqlService.PostCreateUser(req.User); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MySQLRouter) createContent(c *gin.Context) {
	var req CreateContentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mysqlService.PostCreateContent(req.Content, req.Price); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MySQLRouter) buy(c *gin.Context) {
	var req BuyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mysqlService.PostBuy(req.User); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MySQLRouter) userBucket(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mysqlService.GetUserBucket(req.User); err != nil {
		if err.Error() == NoSQLResult {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}

}

func (m *MySQLRouter) content(c *gin.Context) {
	var req ContentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mysqlService.GetContent(req.Content); err != nil {
		if err.Error() == NoSQLResult {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MySQLRouter) userBucketHistory(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mysqlService.GetUserHistory(req.User); err != nil {
		if err.Error() == NoSQLResult {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MySQLRouter) health(c *gin.Context) {
	time.Sleep(3 * time.Second)

	if c.Request.Context().Err() == context.DeadlineExceeded {
		fmt.Println("에러가 났습니다.")
	} else {
		fmt.Println("들어옵니다.")
	}

	if !c.Writer.Written() {
		c.JSON(200, "test")
	}
}
