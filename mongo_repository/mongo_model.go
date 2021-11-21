package mongo_repository

import "time"

type ArtworkMongo struct {
	ArtworkID    string      `bson:"_id"`
	ArtistUserID string      `bson:"artist_id"`
	Title        string      `bson:"art_title"`
	Description  string      `bson:"art_desc"`
	UploadDate   time.Time   `bson:"upload_date"`
	ArtworkUrl   string      `bson:"art_url"`
	Tags         []TagsMongo `bson:"tags"`
}
type TagsMongo struct {
	TagID   string `bson:"_id"`
	TagName string `bson:"tag_name"`
}
