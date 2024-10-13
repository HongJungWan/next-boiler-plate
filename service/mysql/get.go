package mysql

import (
	"eCommerce/types"
	"fmt"
)

func (m *MySQLService) GetUserBucket(user string) (*types.User, error) {
	if r, err := m.repository.MySQL.GetUserBucket(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}

func (m *MySQLService) GetContent(name string) (*types.ContentResponse, error) {
	if r, err := m.repository.MySQL.GetContent(name); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return &types.ContentResponse{
			ContentList: r,
			Description: "와 성공이다!!",
			ResultCode:  1,
		}, err
	}
}

func (m *MySQLService) GetUserHistory(user string) (*types.History, error) {
	if r, err := m.repository.MySQL.GetUserHistory(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}
