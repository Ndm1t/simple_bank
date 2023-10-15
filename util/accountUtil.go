package util

import "math"

var CurrMap = map[string]int64{
	"EUR": 978,
	"USD": 840,
	"UAH": 980,
}

const EURtoUSD = 1.05

func ConvertBalance(currFrom string, balance float64) (convertedBalance float64) {
	if currFrom == "USD" {
		convertedBalance = balance / EURtoUSD
	} else {
		convertedBalance = balance * EURtoUSD
	}
	return math.Round(convertedBalance*100) / 100
}

/*type CurrData struct {
	CurrencyCodeA int64   `json:"currencyCodeA"`
	CurrencyCodeB int64   `json:"currencyCodeB"`
	Date          int64   `json:"date"`
	RateSell      float64 `json:"rateSell"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
}

func ConvertBalance(currFrom string, currTo string, balance float64) (convertedBalance float64, err error) {
	res, err := http.Get("https://api.monobank.ua/bank/currency")

	if res.StatusCode == http.StatusTooManyRequests {
		err = errors.New("too many requests")
		return
	}

	if err != nil {
		return
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var CurrArr []CurrData

	if err = json.Unmarshal(responseData, &CurrArr); err != nil {
		return
	}

	var rateSell, rateBuy float64

	for _, v := range CurrArr {
		if v.CurrencyCodeA == CurrMap[currFrom] && v.CurrencyCodeB == CurrMap["UAH"] {
			rateSell = v.RateSell
		} else if v.CurrencyCodeA == CurrMap[currTo] && v.CurrencyCodeB == CurrMap["UAH"] {
			rateBuy = v.RateBuy
		}
	}

	newBalance := (balance * rateSell) / rateBuy
	convertedBalance = math.Round(newBalance*100) / 100

	return
}
*/
