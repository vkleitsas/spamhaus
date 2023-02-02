package app

import (
	"assignment/mocks"
	"assignment/testdata"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// valid url test
func TestSaveWebsiteURL_OK(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)
	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}

	urlDownload.EXPECT().DownloadSingleUrl(gomock.Any()).Return(&testdata.SaveWebsiteURLTestsData[0].Entry, nil)
	repo.EXPECT().Update(gomock.Any()).Return(&testdata.SaveWebsiteURLTestsData[0].Entry, nil)

	res, err := crudService.SaveWebsiteURL(testdata.SaveWebsiteURLTestsData[0].Entry)
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	if !reflect.DeepEqual(res, testdata.SaveWebsiteURLTestsData[0].Want) {
		t.Errorf("SaveWebsiteURL() = %v, want %v", res, testdata.SaveWebsiteURLTestsData[0].Want)
	}

}
func TestSaveWebsiteURL_InvalidURL(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)
	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}

	_, err := crudService.SaveWebsiteURL(testdata.SaveWebsiteURLTestsData[1].Entry)
	if err == nil {
		t.Error("Error expected")
	}

}

func TestSaveWebsiteURL_ErrorfromUrlDownload(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)
	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}
	urlDownload.EXPECT().DownloadSingleUrl(gomock.Any()).Return(nil, errors.New("some error"))
	_, err := crudService.SaveWebsiteURL(testdata.SaveWebsiteURLTestsData[2].Entry)
	if err == nil {
		t.Error("Error expected")
	}

}

func TestSaveWebsiteURL_ErrorfromRepo(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)
	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}
	urlDownload.EXPECT().DownloadSingleUrl(gomock.Any()).Return(&testdata.SaveWebsiteURLTestsData[2].Entry, nil)
	repo.EXPECT().Update(gomock.Any()).Return(nil, errors.New("some error"))

	_, err := crudService.SaveWebsiteURL(testdata.SaveWebsiteURLTestsData[3].Entry)
	if err == nil {
		t.Error("Error expected")
	}

}

func TestGetWebsiteURLS_OK_1(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)

	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}

	repo.EXPECT().GetByDate().Return(testdata.TestGetWebsiteURLS_OK_Data, nil)

	res, err := crudService.GetWebsiteURLS("date")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	if !reflect.DeepEqual(res, testdata.TestGetWebsiteURLS_OK_Data) {
		t.Errorf("SaveWebsiteURL() = %v, want %v", res, testdata.TestGetWebsiteURLS_OK_Data)
	}
}

func TestGetWebsiteURLS_OK_2(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)

	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}

	repo.EXPECT().GetBySize().Return(testdata.TestGetWebsiteURLS_OK_Data, nil)

	res, err := crudService.GetWebsiteURLS("size")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	if !reflect.DeepEqual(res, testdata.TestGetWebsiteURLS_OK_Data) {
		t.Errorf("SaveWebsiteURL() = %v, want %v", res, testdata.TestGetWebsiteURLS_OK_Data)
	}
}

func TestGetWebsiteURLS_OK_InvalidSort(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	urlDownload := mocks.NewMockUrlDownloadInterface(ctrl)
	repo := mocks.NewMockURLRequestInterface(ctrl)

	crudService := &CrudService{
		UrlDownload: urlDownload,
		Repo:        repo,
	}

	_, err := crudService.GetWebsiteURLS("invalid")
	if err == nil {
		t.Error("Error expected")
	}
	if err.Error() != "sort must be date or size" {
		t.Error("Unexpected error: ", err)
	}

}
