package mongo

import (
	"errors"
	"fmt"
)

func (m *MService) PostCreateUser(user string) error {

	if err := m.repository.Mongo.PostCreateUser(user); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}

}

func (m *MService) PostCreateContent(user string, price int64) error {

	if err := m.repository.Mongo.PostCreateContent(user, price); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}

}

func (m *MService) PostBucketRequest(user, content string) error {
	/*
		1. user가 존재 하는지
		2. content가 존재 하는지
	*/

	if c, err := m.repository.Mongo.GetContent(content); err != nil {
		fmt.Println("GetContent err", err)
		return err
	} else if _, err = m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err", err)
		return err
	} else if len(c) == 0 {
		return errors.New("Content 없습니다.")
	} else if err = m.repository.Mongo.PostInsertBucket(user, content); err != nil {
		fmt.Println("PostInsertBucket err", err)
		return err
	} else {
		return nil
	}

}

func (m *MService) PostBuy(user string) error {

	/*
		1. 기존 유저에 있는 장바구니 데이터를 history에 넣어준다.
	*/

	if u, err := m.repository.Mongo.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err", err)
		return err
	} else if len(u.Bucket) == 0 {
		return errors.New("장바구니에 데이터가 없습니다.")
	} else {
		// -> history컬렉션에 데이터 넣어주기

		if err = m.repository.Mongo.UpsertHistory(user, u.Bucket); err != nil {
			fmt.Println("UpsertHistory err", err)
			return err
		}

		// -> user에서 bucket필드 초기화 하기
		if err = m.repository.Mongo.RemoveUserBucket(user); err != nil {
			fmt.Println("RemoveUserBucket err", err)
			return err
		}
	}

	return nil
}
