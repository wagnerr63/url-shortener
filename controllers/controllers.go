package controllers

import (
	shortenedurl "url-shortener/controllers/shortenedUrl"
	"url-shortener/usecases"
)

type Container struct {
	ShortenedUrl shortenedurl.IShortenedUrlController
}

type Options struct {
	Usecases *usecases.Container
}

func New(opts Options) *Container {
	return &Container{
		ShortenedUrl: shortenedurl.New(opts.Usecases),
	}
}
