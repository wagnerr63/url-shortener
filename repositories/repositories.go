package repositories

import "url-shortener/repositories/shortenedUrls"

type Container struct {
	ShortenedUrl shortenedUrls.IShortnedUrlsRepository
}

// type Options struct {
// 	WriterGorm *gorm.DB
// 	ReaderGorm *gorm.DB
// }

func New() *Container {
	return &Container{
		ShortenedUrl: shortenedUrls.NewMockShortenedUrlsRepository(),
	}
}
