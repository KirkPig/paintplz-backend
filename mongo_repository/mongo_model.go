package mongo_repository

type ArtworkMongo struct {
	ArtworkID   string      `bson:"_id"`
	Title       string      `bson:"art_title"`
	Description string      `bson:"art_desc"`
	UploadDate  string      `bson:"upload_date"`
	ArtworkUrl  string      `bson:"art_url"`
	Tags        []TagsMongo `bson:"tags"`
}
type TagsMongo struct {
	TagID   string `bson:"_id"`
	TagName string `bson:"tag_name"`
}
