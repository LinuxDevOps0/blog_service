package model

import "practice.com/blog_service/pkg/app"

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List   []*Tag
	Pager  *app.Pager
}
