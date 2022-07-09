package banking

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var format = "\n%s - покупка: %.2f"
var NetClient = http.Client{Timeout: time.Second * 1000}

const (
	nationalBankApi = "https://bank.gov.ua/NBUStatService/v1/statdirectory/exchange?json"
	privatbankApiG  = "https://api.privatbank.ua/p24api/pubinfo?json&exchange&coursid=5"
	privatBankApiB  = "https://api.privatbank.ua/p24api/pubinfo?exchange&json&coursid=11"
)

func decode(URL string) []byte {
	res, err := NetClient.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func nationalBank() (msg string) {
	body := decode(nationalBankApi)
	var result []NationalBank
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result {
		switch rec.Cc {
		case "USD":
			msg = fmt.Sprintf("*\nНациональный банк Украины:*"+format, rec.Cc, rec.Rate)
		case "EUR":
			msg += fmt.Sprintf(format, rec.Cc, rec.Rate)
		}
	}
	return
}

func privatBankDLC() (msg string) {
	body := decode(privatbankApiG)

	var result []PrivatBank24
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result {
		switch rec.Ccy {
		case "USD":
			msg = fmt.Sprintf("\nГотівковий курс *PrivatBank:*\n%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "EUR":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "RUR":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "BTC":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		}
	}
	return
}

func Banks() (msg string) {
	body := decode(privatBankApiB)
	var result []PrivatBank24
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	for _, rec := range result {
		switch rec.Ccy {
		case "USD":
			msg += fmt.Sprintf("\nБезготівковий курс *PrivatBank:*\n%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "EUR":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "RUR":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		case "BTC":
			msg += fmt.Sprintf("%s - покупка: %.2f, продажа: %.2f\n", rec.Ccy, rec.Buy, rec.Sale)
		}
	}
	msg += privatBankDLC() + nationalBank()
	return
}

func Weather(city, api string) (msg string) {
	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/find?q=%s,UA&units=metric&lang=RU&type=like&APPID=%s", city, api)

	body := decode(URL)
	var result OpenWeather
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	if len(result.List) <= 0 {
		msg = "такого города нет, иди нахуй"
		return
	}

	msg = fmt.Sprintf("Город: %s, %s\nТемпература: %.2f С\nОписание: %s", result.List[0].Name, result.List[0].Sys.Country, result.List[0].Main.FeelsLike, result.List[0].Weather[0].Description)
	return
}
