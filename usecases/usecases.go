package usecases

import (
	"url-shortener/repositories"
	shortenedurl "url-shortener/usecases/shortenedUrl"
)

type Container struct {
	ShortenedUrl shortenedurl.IShortenedUrlUseCases
}

type Options struct {
	Repo *repositories.Container
}

func New(opts Options) *Container {

	return &Container{
		ShortenedUrl: shortenedurl.New(opts.Repo),
	}
}
