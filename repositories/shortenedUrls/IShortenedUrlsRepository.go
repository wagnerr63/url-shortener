package shortenedUrls

import "url-shortener/entities"

type IShortnedUrlsRepository interface {
	Create(shortenedUrl entities.ShortenedUrl) error
	FindAll() ([]entities.ShortenedUrl, error)
	FindByUrl(url string) (entities.ShortenedUrl, error)
	FindByShortenedID(id string) (entities.ShortenedUrl, error)
}
