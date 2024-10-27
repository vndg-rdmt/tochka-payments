package service

import "time"

type PaymentRequest struct {
	Data PaymentRequestData `json:"Data"`
}

type PaymentRequestData struct {
	CustomerCode    string   `json:"customerCode"`
	Amount          string   `json:"amount"`
	Purpose         string   `json:"purpose"`
	RedirectUrl     string   `json:"redirectUrl"`
	FailRedirectUrl string   `json:"failRedirectUrl"`
	PaymentMode     []string `json:"paymentMode"`
	SaveCard        bool     `json:"saveCard"`
	ConsumnerId     string   `json:"consumerId"`
}

type PaymentResponse struct {
	Data struct {
		OperationID string `json:"operationId"`
	} `json:"Data"`
}

type CheckStatus struct {
	Data struct {
		Operation []struct {
			CustomerCode    string    `json:"customerCode"`
			TaxSystemCode   string    `json:"taxSystemCode"`
			PaymentType     string    `json:"paymentType"`
			PaymentID       string    `json:"paymentId"`
			TransactionID   string    `json:"transactionId"`
			CreatedAt       time.Time `json:"createdAt"`
			PaymentMode     []string  `json:"paymentMode"`
			RedirectURL     string    `json:"redirectUrl"`
			FailRedirectURL string    `json:"failRedirectUrl"`
			Client          struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				Phone string `json:"phone"`
			} `json:"Client"`
			Items []struct {
				VatType       string `json:"vatType"`
				Name          string `json:"name"`
				Amount        string `json:"amount"`
				Quantity      int    `json:"quantity"`
				PaymentMethod string `json:"paymentMethod"`
				PaymentObject string `json:"paymentObject"`
				Measure       string `json:"measure"`
			} `json:"Items"`
			Purpose     string `json:"purpose"`
			Amount      string `json:"amount"`
			Status      string `json:"status"`
			OperationID string `json:"operationId"`
			PaymentLink string `json:"paymentLink"`
			MerchantID  string `json:"merchantId"`
			ConsumerID  string `json:"consumerId"`
		} `json:"Operation"`
	} `json:"Data"`
	Links struct {
		Self string `json:"self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"totalPages"`
	} `json:"Meta"`
}
