package database

import (
	"context"
	"testing"

	"github.com/CayenneLow/Codenames/internal/config"
	"github.com/CayenneLow/Codenames/internal/game"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	db  Database
	ctx context.Context
}

func (suite *TestSuite) SetupTest() {
	ctx := context.Background()
	db := Init(ctx, config.Init())
	suite.ctx = ctx
	suite.db = db
}

func (suite *TestSuite) TestInsert() {
	mockGameID := "INTT3"
	mockData := `{"cells":[[{"word":"1x1","team":"BLUE","guessed":false},{"word":"1x2","team":"BLUE","guessed":false}],[{"word":"2x1","team":"RED","guessed":false},{"word":"2x2","team":"RED","guessed":false}]]}`
	mockDataStruct, err := game.FromJson([]byte(mockData))
	assert.NoError(suite.T(), err)
	err = suite.db.Insert(suite.ctx, mockGameID, mockDataStruct)
	assert.NoError(suite.T(), err)

	actualBoard, err := suite.db.ReadBoardByGameID(suite.ctx, mockGameID)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockDataStruct, actualBoard)
}

func (suite *TestSuite) TestReadBoardByGameID() {
	expectedJson := `{"cells":[[{"word":"1x1","team":"BLUE","guessed":false},{"word":"1x2","team":"RED","guessed":false}],[{"word":"2x1","team":"BLUE","guessed":false},{"word":"2x2","team":"RED","guessed":false}]]}`
	expectedDataStruct, err := game.FromJson([]byte(expectedJson))
	assert.NoError(suite.T(), err)
	actual, err := suite.db.ReadBoardByGameID(suite.ctx, "INTT1")
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedDataStruct, actual)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
