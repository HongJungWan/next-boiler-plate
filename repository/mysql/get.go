package mysql

import (
	"context"
	"eCommerce/types"
	. "eCommerce/types/err"
	"encoding/json"
	"errors"
)

// queryContext : mongo의 find -> 여러개의 데이터를 가져오는 경우 사용
// queryRowContext : mongo의 FindOne -> 한개의 데이터를 가져오는 경우

func (m *MySQL) GetUser(user string) (*types.User, error) {
	u := new(types.User)

	err := m.DB.QueryRowContext(
		context.Background(), "SELECT * FROM eCommerce.user WHERE user = ?;", user).
		Scan(&u.User)

	return u, err
}

func (m *MySQL) GetUserBucket(user string) (*types.User, error) {
	u := new(types.User)

	var jsonParsing []uint8

	err := m.DB.QueryRowContext(
		context.Background(), "select * from eCommerce.bucket WHERE user = ?;", user).
		Scan(&u.User, &jsonParsing)

	if err == nil {
		var bucket []string
		if err = json.Unmarshal(jsonParsing, &bucket); err != nil {
			return nil, err
		} else {
			u.Bucket = bucket
		}
	}

	return u, err
}

func (m *MySQL) GetContent(name string) ([]*types.Content, error) {

	ctx := context.Background()

	if name == "" {
		cursor, err := m.DB.QueryContext(ctx, "SELECT name, price FROM eCommerce.content;")

		if err != nil {
			return nil, err
		} else {
			defer cursor.Close()

			list := make([]*types.Content, 0)

			for cursor.Next() {
				c := new(types.Content)

				if err = cursor.Scan(
					&c.Name,
					&c.Price,
				); err != nil {
					return nil, err
				} else {
					list = append(list, c)
				}
			}

			if len(list) == 0 {
				return nil, errors.New(NoSQLResult)
			} else {
				return list, nil
			}
		}
		// 전체 가져오기
	} else {
		// 있다면, 해당 값만 매칭해서 가져오기
		c := new(types.Content)

		if err := m.DB.QueryRowContext(
			context.Background(), "SELECT name, price FROM eCommerce.content WHERE name = ?;", name).
			Scan(&c.Name, &c.Price); err != nil {
			return nil, err
		} else {
			return []*types.Content{c}, err
		}

	}
}

func (m *MySQL) GetUserHistory(user string) (*types.History, error) {
	u := new(types.History)

	var jsonParsing []uint8

	err := m.DB.QueryRowContext(
		context.Background(), "SELECT user, contentList FROM eCommerce.history WHERE user = ?;", user).
		Scan(&u.User, &jsonParsing)

	if err == nil {
		var list []string
		if err = json.Unmarshal(jsonParsing, &list); err != nil {
			return nil, err
		} else {
			u.ContentList = list
		}
	}

	return u, err
}
