package shortenedUrls

import (
	"errors"
	"url-shortener/entities"
)

type MockShortenedUrlsRepository struct {
	shortenedUrls []entities.ShortenedUrl
}

func NewMockShortenedUrlsRepository() IShortnedUrlsRepository {
	return &MockShortenedUrlsRepository{shortenedUrls: nil}
}

func (repo *MockShortenedUrlsRepository) Create(shortenedUrl entities.ShortenedUrl) error {
	repo.shortenedUrls = append(repo.shortenedUrls, shortenedUrl)
	return nil
}

func (repo *MockShortenedUrlsRepository) FindAll() ([]entities.ShortenedUrl, error) {
	return repo.shortenedUrls, nil
}

func (repo *MockShortenedUrlsRepository) FindByUrl(url string) (entities.ShortenedUrl, error) {
	for _, shortenedUrl := range repo.shortenedUrls {
		if shortenedUrl.Url == url {
			return shortenedUrl, nil
		}
	}
	return entities.ShortenedUrl{}, errors.New("shortened url not found")
}

func (repo *MockShortenedUrlsRepository) FindByShortenedID(id string) (entities.ShortenedUrl, error) {
	for _, shortenedUrl := range repo.shortenedUrls {
		if shortenedUrl.ShortenedID == id {
			return shortenedUrl, nil
		}
	}
	return entities.ShortenedUrl{}, errors.New("shortened url not found")
}
