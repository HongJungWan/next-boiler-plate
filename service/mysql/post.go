package mysql

import (
	err2 "eCommerce/types/err"
	"errors"
	"fmt"
)

func (m *MySQLService) PostCreateUser(user string) error {

	if err := m.repository.MySQL.PostCreateUser(user); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}

}

func (m *MySQLService) PostCreateContent(content string, price int64) error {

	if err := m.repository.MySQL.PostCreateContent(content, price); err != nil {
		fmt.Println(err)
		return err
	} else {
		return err
	}

}

func (m *MySQLService) PostBucketRequest(user, content string) error {
	/*
		1. user가 존재 하는지
		2. content가 존재 하는지
	*/

	if c, err := m.repository.MySQL.GetContent(content); err != nil {
		fmt.Println("GetContent err", err)
		return err
	} else if _, err := m.repository.MySQL.GetUser(user); err != nil {
		fmt.Println("GetUser err", err)
		return err
	} else if len(c) == 0 {
		return errors.New("Content 없습니다.")
	} else if b, err := m.repository.MySQL.GetUserBucket(user); err != nil {
		if err.Error() == err2.NoSQLResult {
			// 처음으로 장바구니에 추가하는 유저
			if err = m.repository.MySQL.PostInsertBucket(user, []string{content}); err != nil {
				fmt.Println("PostInsertBucket err", err)
				return err
			} else {
				return nil
			}
		} else {
			fmt.Println("GetUserBucket err", err)
			return err
		}
		// 에러가 발생을 한다면, 처음 사용자
	} else {
		b.Bucket = append(b.Bucket, content)

		if err = m.repository.MySQL.UpdateBucket(user, b.Bucket); err != nil {
			fmt.Println("UpdateBucket err", err)
			return err
		} else {
			return nil
		}

	}

}

func (m *MySQLService) PostBuy(user string) error {

	/*
		1. 기존 유저에 있는 장바구니 데이터를 history에 넣어준다.
	*/

	if u, err := m.repository.MySQL.GetUserBucket(user); err != nil {
		fmt.Println("GetUserBucket err", err)
		return err
	} else if len(u.Bucket) == 0 {
		return errors.New("장바구니에 데이터가 없습니다.")
	} else {
		// -> history컬렉션에 데이터 넣어주기

		if h, err := m.repository.MySQL.GetUserHistory(user); err != nil {
			if err.Error() == err2.NoSQLResult {
				// insert 쿼리 실행
				if err = m.repository.MySQL.InsertUserHistory(user, u.Bucket); err != nil {
					fmt.Println("InsertUserHistory err", err)
					return err
				}
			} else {
				fmt.Println("GetUserHistory err", err)
				return err
			}
		} else {
			u.Bucket = append(u.Bucket, h.ContentList...)
			// update 쿼리 실행
			if err = m.repository.MySQL.UpdateUserHistory(user, u.Bucket); err != nil {
				fmt.Println("InsertUserHistory err", err)
				return err
			}
		}

		// -> user에서 bucket필드 초기화 하기

		if err := m.repository.MySQL.RemoveUserBucket(user); err != nil {
			fmt.Println("RemoveUserBucket err", err)
			return err
		}
	}

	return nil
}
