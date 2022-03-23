package repositories

import (
	"url-shortener/repositories/shortenedUrls"

	"gorm.io/gorm"
)

type Container struct {
	ShortenedUrl shortenedUrls.IShortnedUrlsRepository
}

type Options struct {
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

func New(opts Options) *Container {
	return &Container{
		ShortenedUrl: shortenedUrls.NewGormRepository(opts.WriterGorm, opts.ReaderGorm),
	}
}
