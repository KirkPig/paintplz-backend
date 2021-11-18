package services

type UploadArtworkRequest struct {
	Username           string `json:"username" binding:"required"`
	ArtworkName        string `json:"artworkName" binding:"required"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl" binding:"required"`
}

type UploadArtworkResponse struct {
	Log string `json:"log"`
}

type EditArtworkRequest struct {
	Username           string `json:"username" binding:"required"`
	ArtworkID          int    `json:"artworkID" binding:"required"`
	ArtworkName        string `json:"artworkName"`
	ArtworkDescription string `json:"artworkDescription"`
	ArtTag             []Tag  `json:"artTag"`
	ArtworkUrl         string `json:"artworkUrl"`
}

type EditArtworkResponse struct {
	Log string `json:"log"`
}

type DeleteArtworkRequest struct {
	Username  string `json:"username" binding:"required"`
	ArtworkID int    `json:"artworkID" binding:"required"`
}

type DeleteArtworkResponse struct {
	Log string `json:"log"`
}

type GetTagResponse struct {
	Tags []Tag `json:"tags"`
}

type Tag struct {
	TagId   string `json:"tagID" binding:"required"`
	TagName string `json:"tagName" binding:"required"`
}
