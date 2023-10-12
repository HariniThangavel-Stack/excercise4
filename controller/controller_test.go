package controller

import (
	"context"
	"excercise4/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct{}

func (m *mockCollection) Find(ctx context.Context, filter interface{},
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	c := &mongo.Cursor{}
	return c, nil
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

func (m *mockCollection) FindOneAndUpdate(ctx context.Context, filter interface{},
	update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	c := &mongo.SingleResult{}
	return c

}
func (m *mockCollection) ReplaceOne(ctx context.Context, filter interface{},
	replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	c := &mongo.UpdateResult{}
	return c, nil
}

func TestBase(t *testing.T) {
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(Base)
	handler.ServeHTTP(res, req)
	assert.Equal(t, "Server Listening on port 3000", res.Body.String(), "Incorrect body found")
}

func TestGetData(t *testing.T) {
	mockCol := &mockCollection{}
	res, err := GetData(mockCol, "BANDHANBNK123")
	assert.Nil(t, err)
	assert.IsType(t, &mongo.Cursor{}, res)
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	data := model.StockMarketData{
		Symbol:                "BANDHANBNK123",
		Open:                  "302.90",
		High:                  "310.95",
		Low:                   "298.20",
		PreviousClose:         "301.95",
		Ltp:                   "307.25",
		Change:                "5.30",
		PercentageChange:      "1.76",
		Volume:                "8859238",
		Value:                 "2,700,029,965.26",
		YearHigh:              "430.70",
		YearLow:               "251.40",
		YearPercentageChange:  "-1.61",
		MonthPercentageChange: "-6.39",
	}
	res, err := InsertData(mockCol, data)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}

func TestDeleteData(t *testing.T) {
	mockCol := &mockCollection{}
	res, err := DeleteData(mockCol, "BANDHANBNK123")
	assert.Nil(t, err)
	assert.IsType(t, &mongo.DeleteResult{}, res)
}

func TestUpdateData(t *testing.T) {
	mockCol := &mockCollection{}
	data := model.UpdateBody{
		Symbol: "BANDHANBNK",
	}
	res := UpdateData(mockCol, "BANDHANBNK123", data)
	assert.IsType(t, &mongo.SingleResult{}, res)
}

func TestReplaceOneDoc(t *testing.T) {
	mockCol := &mockCollection{}
	data := model.StockMarketData{
		Symbol:                "BANDHANBNK123",
		Open:                  "302.90",
		High:                  "310.95",
		Low:                   "298.20",
		PreviousClose:         "301.95",
		Ltp:                   "307.25",
		Change:                "5.30",
		PercentageChange:      "1.76",
		Volume:                "8859238",
		Value:                 "2,700,029,965.26",
		YearHigh:              "430.70",
		YearLow:               "251.40",
		YearPercentageChange:  "-1.61",
		MonthPercentageChange: "-6.39",
	}
	res, err := ReplaceOneDoc(mockCol, "BANDHANBNK", data)
	assert.Nil(t, err)
	assert.IsType(t, &mongo.UpdateResult{}, res)
}
