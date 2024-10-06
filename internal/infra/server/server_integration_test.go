//go:build integration_tests
// +build integration_tests

package server

import (
	"bytes"
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/config"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/db"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/infra/testfun"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AppTestSuite struct {
	suite.Suite
	stopServices context.CancelFunc
}

func (suite *AppTestSuite) SetupSuite() {
	os.Setenv("RUN_ADDRESS", "localhost:8085")
	ctx, cancel := context.WithCancel(context.Background())
	dsn, err := testfun.CreateTestDB(ctx)
	if err != nil {
		panic(err)
	}
	os.Setenv("DATABASE_URI", dsn)
	suite.stopServices = cancel

	config.InitConfig()
	err = db.InitDB(config.AppConfig.Database)
	if err != nil {
		panic(err)
	}
	go func(ctx context.Context) {
		Start(ctx)
	}(ctx)
}

func (suite *AppTestSuite) TestRegistration() {
	t := suite.T()
	resp, err := http.Post("http://localhost:8085/api/user/register", "application/json", bytes.NewBuffer([]byte(`{"login": "sdf","password":"qwer"}`)))
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, 200)
	assert.NotEmpty(t, resp.Header.Get("Authorization"))
}

func (suite *AppTestSuite) TearDownSuite() {
	suite.stopServices()
	time.Sleep(time.Second) // FIXME: shutdown tests
}

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
