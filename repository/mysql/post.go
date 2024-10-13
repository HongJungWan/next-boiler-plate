package mysql

import (
	"context"
	"encoding/json"
)

// ExecContext

func (m *MySQL) PostCreateUser(user string) error {
	_, err := m.DB.ExecContext(context.Background(), "INSERT INTO eCommerce.user(user) VALUES(?)", user)
	return err
}

func (m *MySQL) PostCreateContent(name string, price int64) error {
	_, err := m.DB.ExecContext(context.Background(), "INSERT INTO eCommerce.content(name, price) VALUES(?, ?)", name, price)
	return err
}

func (m *MySQL) PostInsertBucket(user string, content []string) error {
	// 배열에 값을 추가해줘야 합니다.
	j, _ := json.Marshal(content)
	_, err := m.DB.ExecContext(context.Background(), "INSERT INTO eCommerce.bucket(user, bucket) VALUES(?, ?)", user, j)
	return err
}

func (m *MySQL) UpdateBucket(user string, content []string) error {
	j, _ := json.Marshal(content)
	_, err := m.DB.ExecContext(context.Background(), "UPDATE eCommerce.bucket SET bucket = ? WHERE user = ?;", j, user)
	return err
}

func (m *MySQL) InsertUserHistory(user string, contentList []string) error {
	j, _ := json.Marshal(contentList)
	_, err := m.DB.ExecContext(context.Background(), "INSERT INTO eCommerce.history(user, contentList) VALUES(?, ?)", user, j)
	return err
}

func (m *MySQL) UpdateUserHistory(user string, contentList []string) error {
	j, _ := json.Marshal(contentList)
	_, err := m.DB.ExecContext(context.Background(), "UPDATE eCommerce.history SET contentList = ? WHERE user = ?;", j, user)
	return err
}

func (m *MySQL) RemoveUserBucket(user string) error {
	_, err := m.DB.ExecContext(context.Background(), "DELETE FROM eCommerce.bucket WHERE user = ?;", user)
	return err
}
