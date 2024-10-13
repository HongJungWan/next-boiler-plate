package mongo

import (
	"context"
	"eCommerce/types"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongo) GetUserBucket(user string) (*types.User, error) {
	filter := bson.M{"user": user}

	var u types.User

	if err := m.user.FindOne(context.Background(), filter).Decode(&u); err != nil {
		return nil, err
	} else {
		return &u, nil
	}

}

func (m *Mongo) GetContent(name string) ([]*types.Content, error) {

	filter := bson.M{}

	if name != "" {
		filter["name"] = name
	}

	ctx := context.Background()

	if cursor, err := m.content.Find(ctx, filter); err != nil {
		return nil, err
	} else {
		defer cursor.Close(ctx)

		var v []*types.Content

		if err = cursor.All(ctx, &v); err != nil {
			return nil, err
		} else {
			return v, nil
		}

	}
}

func (m *Mongo) GetUserHistory(user string) (*types.History, error) {
	filter := bson.M{"user": user}

	var h types.History

	if err := m.history.FindOne(context.Background(), filter).Decode(&h); err != nil {
		return nil, err
	} else {
		return &h, nil
	}

}
