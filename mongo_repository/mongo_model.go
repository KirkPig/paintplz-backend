package mongo_repository

import "time"

type ArtworkMongo struct {
	ArtworkID    string      `bson:"_id" json:"artworkID"`
	ArtistUserID string      `bson:"artist_id" json:"artistID"`
	Title        string      `bson:"art_title" json:"artTitle"`
	Description  string      `bson:"art_desc" json:"artDescription"`
	UploadDate   time.Time   `bson:"upload_date" json:"uploadDate"`
	ArtworkUrl   string      `bson:"art_url" json:"artworkUrl"`
	Tags         []TagsMongo `bson:"tags" json:"tags"`
}
type TagsMongo struct {
	TagID   string `bson:"_id" json:"tagID"`
	TagName string `bson:"tag_name" json:"tagName"`
}
