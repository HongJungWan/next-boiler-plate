package mongo

import (
	"eCommerce/types"
	"fmt"
)

func (m *MService) GetUserBucket(user string) (*types.User, error) {
	if r, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}

func (m *MService) GetContent(name string) ([]*types.Content, error) {
	if r, err := m.repository.Mongo.GetContent(name); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}

func (m *MService) GetUserHistory(user string) (*types.History, error) {
	if r, err := m.repository.Mongo.GetUserHistory(user); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return r, err
	}
}
