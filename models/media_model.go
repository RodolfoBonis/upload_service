package models

import "mime/multipart"

type FileModel struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type UrlModel struct {
	Url string `json:"url,omitempty" validate:"required"`
}
