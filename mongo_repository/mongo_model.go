package mongo_repository

import "time"

type ArtworkMongo struct {
	ArtworkID    string      `bson:"_id" json:"artworkID"`
	ArtistUserID string      `bson:"artist_id"`
	Title        string      `bson:"art_title" json:"title"`
	Description  string      `bson:"art_desc" json:"description"`
	UploadDate   time.Time   `bson:"upload_date" json:"uploadDate"`
	ArtworkUrl   string      `bson:"art_url" json:"url"`
	Tags         []TagsMongo `bson:"tags" json:"artTag"`
}
type TagsMongo struct {
	TagID   string `bson:"_id" json:"tagID"`
	TagName string `bson:"tag_name" json:"tagName"`
}
