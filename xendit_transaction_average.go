package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'getUserTransaction' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER uid
 *  2. STRING txnType
 *  3. STRING monthYear
 *
 *  https://jsonmock.hackerrank.com/api/transactions/search?userId=
 */

type LocationDetail struct {
	Id      int32  `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode string `json:"zipCode"`
}

type TransactionData struct {
	Id        int32          `json:"id"`
	UserId    int32          `json:"userId"`
	Timestamp int64          `json:"timestamp"`
	TxnType   string         `json:"txnType"`
	Amount    string         `json:"amount"`
	Location  LocationDetail `json:"location"`
}

type PageResponse struct {
	Page       int32             `json:"page"`
	PerPage    int32             `json:"per_page"`
	TotalPages int32             `json:"total_pages"`
	Data       []TransactionData `json:"data"`
}

const CREDIT string = "credit"
const DEBIT string = "debit"

func GetUserTransaction(uid int32, txnType string, monthYear string) []int32 {
	result := make([]int32, 1)

	debits := []TransactionData{}
	credits := []TransactionData{}

	responseBody := &PageResponse{}
	page := int32(1)
	url := GetTransactionUrl(uid, page)
	resp, err := http.Get(url)

	if err != nil {
		return []int32{-1}
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(responseBody)

	if len(responseBody.Data) > 0 {
		for _, v := range responseBody.Data {

			transDate := time.Unix(0, v.Timestamp*int64(time.Millisecond))
			transMonthYear := transDate.Format("01-2006")
			if transMonthYear == monthYear {
				if strings.EqualFold(v.TxnType, CREDIT) {
					credits = append(credits, v)
				} else if strings.EqualFold(v.TxnType, DEBIT) {
					debits = append(debits, v)
				}
			}

		}
	}

	for p := int32(2); p <= responseBody.TotalPages; p++ {
		url = GetTransactionUrl(uid, p)
		resp, err = http.Get(url)
		json.NewDecoder(resp.Body).Decode(responseBody)

		if len(responseBody.Data) > 0 {
			for _, v := range responseBody.Data {
				transDate := time.Unix(0, v.Timestamp*int64(time.Millisecond))
				transMonthYear := transDate.Format("01-2006")
				if transMonthYear == monthYear {
					if strings.EqualFold(v.TxnType, CREDIT) {
						credits = append(credits, v)
					} else if strings.EqualFold(v.TxnType, DEBIT) {
						debits = append(debits, v)
					}
				}

			}
		}
	}

	//calculate average spending
	averageSpending := countAverageSpending(debits)

	if strings.EqualFold(CREDIT, txnType) {
		return findTransId(credits, averageSpending)
	} else if strings.EqualFold(DEBIT, txnType) {
		return findTransId(debits, averageSpending)
	}

	return result
}

func findTransId(trans []TransactionData, averageSpending float64) []int32 {
	ids := []int32{}
	for _, v := range trans {
		if convertToFloat(v.Amount) > averageSpending {
			ids = append(ids, v.Id)
		}
	}
	return ids
}

func countAverageSpending(debit []TransactionData) float64 {
	var total float64
	for _, trans := range debit {
		amount := convertToFloat(trans.Amount)
		total = total + amount
	}
	return total / float64(len(debit))
}

func convertToFloat(amountInDollar string) float64 {
	//remove dollar sign
	sanitized := strings.Replace(amountInDollar, "$", "", -1)
	//remove coma char
	sanitized = strings.Replace(sanitized, ",", "", -1)
	amount, err := strconv.ParseFloat(sanitized, 64)
	if err != nil {
		return 0
	}
	return amount
}

func GetTransactionUrl(uid int32, page int32) string {
	return fmt.Sprintf("https://jsonmock.hackerrank.com/api/transactions/search?usserId=%d&page=%d", uid, page)
}

// func main() {
// 	getUserTransaction(4, "debit", "02-2019")
// }
