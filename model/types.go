package model1

type StockMarketData struct {
	Symbol                string `json:"symbol"`
	Open                  string `json:"open"`
	High                  string `json:"high"`
	Low                   string `json:"low"`
	PreviousClose         string `json:"previousClose"`
	Ltp                   string `json:"ltp"`
	Change                string `json:"change"`
	PercentageChange      string `json:"percentageChange"`
	Volume                string `json:"volume" `
	Value                 string `json:"value"`
	YearHigh              string `json:"yearHigh"`
	YearLow               string `json:"yearLow"`
	YearPercentageChange  string `json:"yearPercentageChange"`
	MonthPercentageChange string `json:"monthPercentageChange"`
}

type UpdateBody struct {
	Symbol string `json:"symbol"`
}

type Options struct {
	OutputFormat string
	FilePath     string
}
