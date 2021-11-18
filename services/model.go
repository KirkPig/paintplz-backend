package main

type GetTagResponse struct {
	Tags []Tag `json:"tags" binding:"required"`
}

type Tag struct {
	TagId   string `json:"tagID" binding:"required"`
	TagName string `json:"tagName" binding:"required"`
}
