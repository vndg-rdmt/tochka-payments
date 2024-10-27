package service

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
)

type TochkaError struct {
	Code    string `json:"code"`
	ID      string `json:"id"`
	Message string `json:"message"`
	Errors  []struct {
		ErrorCode string `json:"errorCode"`
		Message   string `json:"message"`
		URL       string `json:"url"`
	} `json:"Errors"`
}

func (t *TochkaError) Error() string {
	return t.Message
}

func New(
	customerCode string,
	tochakToken string,
	tochkaUrl string,
	redirectUrl string,
	failRedirectUrl string,
) Service {
	return &serviceimpl{
		customerCode:    customerCode,
		tochakaToken:    tochakToken,
		tochkaUrl:       tochkaUrl,
		redirectUrl:     redirectUrl,
		failRedirectUrl: failRedirectUrl,
	}
}

type serviceimpl struct {
	customerCode    string
	tochakaToken    string
	tochkaUrl       string
	redirectUrl     string
	failRedirectUrl string
}

// Payment implements Service.
func (s *serviceimpl) Payment(ctx context.Context,
	redirectUrl string,
	failRedirectUrl string,
	amount uint64,
	purpose string,
) (string, error) {
	b, err := json.Marshal(PaymentRequest{
		Data: PaymentRequestData{
			CustomerCode:    s.customerCode,
			Amount:          strconv.FormatUint(amount, 10) + ".00",
			Purpose:         purpose,
			RedirectUrl:     s.redirectUrl,
			FailRedirectUrl: s.failRedirectUrl,
			PaymentMode:     []string{"card"},
			SaveCard:        true,
			ConsumnerId:     s.customerCode,
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		s.tochkaUrl+"/acquiring/v1.0/payments",
		bytes.NewReader(b),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.tochakaToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responsePayload, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		var tochkaerr = new(TochkaError)
		if err = json.Unmarshal(responsePayload, tochkaerr); err != nil {
			return "", err
		}
		return "", tochkaerr
	}

	var data PaymentResponse
	if err = json.Unmarshal(responsePayload, &data); err != nil {
		return "", err
	}

	return data.Data.OperationID, nil
}

// Status implements Service.
func (s *serviceimpl) Status(ctx context.Context, id string) ([]byte, error) {

	req, err := http.NewRequest(
		http.MethodGet,
		s.tochkaUrl+"/acquiring/v1.0/payments/"+id,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.tochakaToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responsePayload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responsePayload, nil
}
