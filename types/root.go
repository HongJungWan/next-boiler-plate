package types

// -> schema

type User struct {
	User   string   `json:"user" bson:"user"`
	Bucket []string `json:"bucket" bson:"bucket"`
}

type Content struct {
	Name      string `json:"name" bson:"name"`
	Price     int64  `json:"price" bson:"price"`
	CreatedAt int64  `json:"-" bson:"createdAt"`
	UpdatedAt int64  `json:"-" bson:"updatedAt"`
}

type ContentResponse struct {
	ResultCode  int64      `json:"resultCode"`
	Description string     `json:"description"`
	ContentList []*Content `json:"result"`
}

type History struct {
	User        string   `json:"user" bson:"user"`
	ContentList []string `json:"contentList" bson:"contentList"`
	CreatedAt   int64    `json:"createdAt" bson:"createdAt"`
}

// -> request

type UserRequest struct {
	User string `form:"user" binding:"required"`
}

type ContentRequest struct {
	Content string `form:"content"`
}

// -> request Post

type CreateContentRequest struct {
	Content string `json:"content"`
	Price   int64  `json:"price"`
}

type CreateUserRequest struct {
	User string `json:"user" binding:"required"`
}

type BuyRequest struct {
	User string `json:"user" binding:"required"`
}

type BucketRequest struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"required"`
}
