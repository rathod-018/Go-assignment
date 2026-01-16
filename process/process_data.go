package processor

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

var claims = []map[string]any{}

func LoadData(productId string) (any,error) {
	dataBytes, dataErr :=os.ReadFile("data/data.json")
	claimBytes, claimErr := os.ReadFile("data/claim.json")
	if dataErr != nil || claimErr != nil{
		return nil,fmt.Errorf("Error while reading file")
	}
	var data map[string]any
	var claim map[string]any

	 if err := json.Unmarshal(dataBytes, &data); err != nil {
        return nil, fmt.Errorf("invalid data.json: %w", err)
    }

    if err := json.Unmarshal(claimBytes, &claim); err != nil {
        return nil, fmt.Errorf("invalid claim.json: %w", err)
    }

	productData, err :=findProductById(data,productId)
	
	if err != nil {
        return nil, fmt.Errorf("error while finding product: %w", err)
    }
 
	for key, value := range claim{
		// fmt.Printf("type of value %T\n",value)

		parts := strings.Split(value.(string),", ")

		for _,value :=range parts{
			ProcessData(key, value, productData)
		}
	}

	// fmt.Println(claims)

	// jsonData, err :=json.Marshal(claims)

	// if err != nil{
	// 	panic(err)
	// }

	// return string(jsonData)
	 return claims, nil

}

func findProductById(data map[string]any, productId string)(any, error){
	dataList, ok := data["data"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("data is missing")
	}

	product, ok := dataList["getProduct"].(map[string]any)
	if !ok {
		return  nil, fmt.Errorf("getProduct missing")
	}

	id, ok := product["id"].(string)

		if !ok {
		return  nil, fmt.Errorf("id missing")
	}

	if id != productId{
		return nil ,fmt.Errorf("product not found")
	}

	return  product, nil

}


func ProcessData(key string, claim string, data any) {

	if m, ok := data.(map[string]any); ok{
		for _, value := range m{
			ProcessData(key, claim, value)
		}
		return
	}

	if s, ok := data.([]any); ok{
		for _, item := range s{
			ProcessData(key, claim,item)
		}
		return
	}


	if _,ok := data.(string); ok{
		isContains := strings.Contains(strings.ToLower(data.(string)), strings.ToLower(claim))
		if isContains{
			id := uuid.New().String()

			newClaim :=Claim{
				ID: id,
				ClaimType: key ,
				ClaimValue: claim,
				Status: "IDENTIFIED",
			}

			result := map[string]any{
				"id":newClaim.ID,
				"claimType":newClaim.ClaimType,
				"claimValue":newClaim.ClaimValue,
				"status":newClaim.Status,
			}

			claims =append(claims,result)
			return
		}
	}

}


type Claim struct {
	ID         string	`json:"id"`
	ClaimType  string 	`json:"claimType"`
	ClaimValue string 	`json:"claimValue"`
	Status     string 	`json:"status"`
}

// workflow struct

type WorkFlow struct{
	WorkFlowId	string
	Product		map[string]any
	Status		string
	Claims		map[string]any
}