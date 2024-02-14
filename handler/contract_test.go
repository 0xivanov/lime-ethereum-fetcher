package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_repo "github.com/0xivanov/lime-ethereum-fetcher-go/mocks"
	"github.com/0xivanov/lime-ethereum-fetcher-go/model"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetPersons_Positive(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := gin.Default()
	mockRepo := mock_repo.NewMockContractInterface(ctrl)
	mockRepo.EXPECT().GetPersons(gomock.Any()).Return([]model.PersonInfoEvent{{
		TxHash: "hash",
		Index:  123,
		Name:   "name",
		Age:    123,
	}}, nil)
	subject := &SmartContract{hclog.Default(), nil, mockRepo}
	router.GET("/lime/persons", subject.GetPersons)

	// when
	req, err := http.NewRequest("GET", "/lime/persons", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// then
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestGetPersons_Negative(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	router := gin.Default()
	mockRepo := mock_repo.NewMockContractInterface(ctrl)
	mockRepo.EXPECT().GetPersons(gomock.Any()).Return(nil, errors.New("mocked error")) // Simulate repository error
	subject := &SmartContract{hclog.Default(), nil, mockRepo}
	router.GET("/lime/persons", subject.GetPersons)

	// when
	req, err := http.NewRequest("GET", "/lime/persons", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// then
	assert.Equal(t, http.StatusInternalServerError, rr.Code) // Expecting Internal Server Error
}
