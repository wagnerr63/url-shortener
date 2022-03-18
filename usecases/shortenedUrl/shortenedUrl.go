package shortenedurl

import (
	"errors"
	"os"
	"strings"
	"time"
	"url-shortener/entities"
	"url-shortener/repositories"
	urlhashid "url-shortener/utils/urlHashID"
)

type ICreateShortenedUrlDTO struct {
	Url string `json:"url"`
}

type ICreateShortenedUrlResponseDTO struct {
	NewUrl string `json:"newUrl"`
}

type IShortenedUrlUseCases interface {
	Create(data ICreateShortenedUrlDTO) (ICreateShortenedUrlResponseDTO, error)
	FindByShortenedID(shortenedID string) (entities.ShortenedUrl, error)
}

type shortenedUrlUseCases struct {
	repositories *repositories.Container
}

func New(repo *repositories.Container) IShortenedUrlUseCases {
	return &shortenedUrlUseCases{
		repositories: repo,
	}
}

func (usecase *shortenedUrlUseCases) Create(data ICreateShortenedUrlDTO) (ICreateShortenedUrlResponseDTO, error) {
	treatedUrl := strings.Replace(data.Url, "https://", "", 1)
	treatedUrl = strings.Replace(treatedUrl, "http://", "", 1)
	treatedUrl = strings.Replace(treatedUrl, "www.", "", 1)

	shortenedUrlFindByUrl, err := usecase.repositories.ShortenedUrl.FindByUrl(treatedUrl)

	if err == nil {
		return ICreateShortenedUrlResponseDTO{
			NewUrl: os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + shortenedUrlFindByUrl.ShortenedID,
		}, nil
	}

	newShortenedUrl := entities.NewShortenedUrl()
	newShortenedUrl.Url = treatedUrl
	newShortenedUrl.ShortenedID = urlhashid.GenerateHash()

	err = usecase.repositories.ShortenedUrl.Create(newShortenedUrl)
	if err != nil {
		return ICreateShortenedUrlResponseDTO{}, errors.New("error to create shortened url")
	}

	return ICreateShortenedUrlResponseDTO{
		NewUrl: os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/" + newShortenedUrl.ShortenedID,
	}, nil
}

func (usecase *shortenedUrlUseCases) FindByShortenedID(shortenedID string) (entities.ShortenedUrl, error) {
	shortenedUrlByShortenedID, err := usecase.repositories.ShortenedUrl.FindByShortenedID(shortenedID)

	if err != nil {
		return entities.ShortenedUrl{}, errors.New("not found")
	}

	parsedExpireDate, _ := time.Parse("2006-01-02 15:04:05", shortenedUrlByShortenedID.ExpiresIn)
	if parsedExpireDate.Unix() < time.Now().Unix() {
		return entities.ShortenedUrl{}, errors.New("shortened url expired")
	}

	return shortenedUrlByShortenedID, nil
}
