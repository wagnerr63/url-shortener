package shortenedUrls

import (
	"errors"
	"fmt"
	"url-shortener/entities"

	"gorm.io/gorm"
)

type repoGorm struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepository(writer, reader *gorm.DB) IShortnedUrlsRepository {
	return &repoGorm{writer: writer, reader: reader}
}

func (repo *repoGorm) Create(shortenedUrl entities.ShortenedUrl) error {
	fmt.Println("aaaaa")
	repo.writer.Table("shortened_urls").Create(&shortenedUrl)

	return repo.writer.Error
}

func (repo *repoGorm) FindAll() ([]entities.ShortenedUrl, error) {
	var shortenedUrls []entities.ShortenedUrl
	repo.reader.Table("shortened_urls").Find(&shortenedUrls)

	if repo.reader.Error != nil {
		return nil, errors.New("shortened urls not found")
	}

	return shortenedUrls, nil
}

func (repo *repoGorm) FindByShortenedID(id string) (entities.ShortenedUrl, error) {
	var shortenedUrlById entities.ShortenedUrl

	repo.reader.Table("shortened_urls").Find(&shortenedUrlById)

	if repo.reader.Error != nil {
		return entities.ShortenedUrl{}, errors.New("shortened url not found")
	}

	return shortenedUrlById, nil
}

func (repo *repoGorm) FindByUrl(url string) (entities.ShortenedUrl, error) {
	var shortenedUrlByUrl entities.ShortenedUrl

	repo.reader.Table("shortened_urls").Where("url = ?", url).Find(&shortenedUrlByUrl)

	if repo.reader.Error != nil {
		return entities.ShortenedUrl{}, errors.New("shortened url not found")
	}

	if (entities.ShortenedUrl{} == shortenedUrlByUrl) {
		return entities.ShortenedUrl{}, errors.New("shortened url not found")
	}

	return shortenedUrlByUrl, nil
}
