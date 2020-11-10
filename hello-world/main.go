package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	zipcode := request.QueryStringParameters["zipcode"]
	if len(zipcode) != 7 	 {
		errorResponse, _ := json.Marshal(&ErrorResponse{
			Message: "Wrong zipcode received",
		})
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: string(errorResponse),
		}, nil
	}
	pref := "東京都"
	city := "品川区"
	town := "五反田"
	address := &AddressResponse{
		Zipcode: zipcode,
		Pref: pref,
		City: city,
		Town: town,
	}

	body, err := json.Marshal(address)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}   

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type AddressResponse struct {
	Zipcode string `json:"zipcode"`
	Pref    string `json:"pref"`
	City    string `json:"city"`
	Town    string `json:"address"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}