package app

import (
	"assignment/domain"
	"assignment/utils"
	"errors"
)

type CrudService struct {
	Repo        domain.URLRequestInterface
	UrlDownload utils.UrlDownloadInterface
}

func NewCrudService(r domain.URLRequestInterface, u utils.UrlDownloadInterface) *CrudService {
	return &CrudService{
		Repo:        r,
		UrlDownload: u,
	}
}

func (c *CrudService) SaveWebsiteURL(r domain.URLEntry) (*domain.URLEntry, error) {
	if r.WebsiteURL == "" {
		return nil, errors.New("Website URL cannot be empty")
	}

	httpResponse, err := c.UrlDownload.DownloadSingleUrl(r)
	if err != nil {
		return nil, err
	}

	result, err := c.Repo.Update(*httpResponse)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *CrudService) GetWebsiteURLS(sort string) ([]domain.URLEntry, error) {
	if sort == "date" {
		result, err := c.Repo.GetByDate()
		if err != nil {
			return nil, err
		}
		return result, nil
	} else if sort == "size" {

		result, err := c.Repo.GetBySize()
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, errors.New("sort must be date or size")
}
