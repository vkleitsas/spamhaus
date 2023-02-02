package testdata

import (
	"assignment/domain"
)

type TestData struct {
	Name  string
	Entry domain.URLEntry
	Want  *domain.URLEntry
}

var SaveWebsiteURLTestsData = []TestData{

	{
		Name: "Valid URL",
		Entry: domain.URLEntry{
			WebsiteURL: "https://www.example.com",
		},
		Want: &domain.URLEntry{
			WebsiteURL: "https://www.example.com",
		},
	},
	{
		Name: "Invalid URL",
		Entry: domain.URLEntry{
			WebsiteURL: "",
		},
		Want: nil,
	},
	{
		Name: "Error from UrlDownload",
		Entry: domain.URLEntry{
			WebsiteURL: "https://www.error.com",
		},
		Want: nil,
	},
	{
		Name: "Error from Repo",
		Entry: domain.URLEntry{
			WebsiteURL: "https://www.example.com",
		},
		Want: nil,
	},
}

var TestGetWebsiteURLS_OK_Data = []domain.URLEntry{{"https://example1.com", 1, []byte("dummy data")}, {"https://example2.com", 2, []byte("dummy data")}}
