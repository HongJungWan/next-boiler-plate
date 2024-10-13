package router

import (
	"context"
	"eCommerce/service/mongo"
	. "eCommerce/types"
	. "eCommerce/types/err"
	"fmt"
	"github.com/gin-gonic/gin"
	pM "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRouter struct {
	router   *Router
	mService *mongo.MService
}

func NewMongoRouter(router *Router, mService *mongo.MService) {
	m := &MongoRouter{
		router:   router,
		mService: mService,
	}

	baseUri := "/mongo"

	m.router.GET(baseUri+"/health", m.health)

	m.router.GET(baseUri+"/user-bucket", m.userBucket)                // 장바구니에 대한 정보를 가져오기
	m.router.GET(baseUri+"/content", m.content)                       // 상품 정보를 가져오는 정보
	m.router.GET(baseUri+"/user-bucket-history", m.userBucketHistory) // 유저의 구매 이력 정보

	m.router.POST(baseUri+"/create-user", m.createUser)       // 유저 데이터 생성
	m.router.POST(baseUri+"/create-content", m.createContent) // content 데이터 생성
	m.router.POST(baseUri+"/buy", m.buy)                      // history 데이터 생성
	m.router.POST(baseUri+"/bucket", m.bucket)                // 장바구니 데이터 생성
}

func (m *MongoRouter) bucket(c *gin.Context) {
	var req BucketRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mService.PostBucketRequest(req.User, req.Content); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MongoRouter) createUser(c *gin.Context) {
	var req CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mService.PostCreateUser(req.User); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MongoRouter) createContent(c *gin.Context) {
	var req CreateContentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mService.PostCreateContent(req.Content, req.Price); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MongoRouter) buy(c *gin.Context) {
	var req BuyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if err = m.mService.PostBuy(req.User); err != nil {
		m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		return
	} else {
		m.router.ResponseOK(c, "Success")
	}
}

func (m *MongoRouter) userBucket(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetUserBucket(req.User); err != nil {
		if err == pM.ErrNoDocuments {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}

}

func (m *MongoRouter) content(c *gin.Context) {
	var req ContentRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetContent(req.Content); err != nil {
		if err == pM.ErrNoDocuments {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MongoRouter) userBucketHistory(c *gin.Context) {
	var req UserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		m.router.ResponseErr(c, ErrorMsg(BindingFailed, err))
		return
	} else if r, err := m.mService.GetUserHistory(req.User); err != nil {
		if err == pM.ErrNoDocuments {
			m.router.ResponseErr(c, ErrorMsg(ServerErr, err))
		} else {
			m.router.ResponseErr(c, ErrorMsg(NoDocument, err))
		}
		return
	} else {
		m.router.ResponseOK(c, r)
	}
}

func (m *MongoRouter) health(c *gin.Context) {
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
