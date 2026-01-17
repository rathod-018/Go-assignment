package process

import (
	"encoding/json"
	"fmt"
	"os"
)


func LoadData(productId string) (any, any, error) {
	dataBytes, dataErr :=os.ReadFile("data/data.json")
	claimBytes, claimErr := os.ReadFile("data/claim.json")
	if dataErr != nil || claimErr != nil{
		return nil, nil, fmt.Errorf("Error while reading file")
	}
	var data map[string]any
	var claim map[string]string

	 if err := json.Unmarshal(dataBytes, &data); err != nil {
        return nil, nil, fmt.Errorf("invalid data.json: %w", err)
    }

    if err := json.Unmarshal(claimBytes, &claim); err != nil {
        return nil, nil, fmt.Errorf("invalid claim.json: %w", err)
    }

	productData, err :=findProductById(data,productId)
	
	if err != nil {
        return nil, nil, fmt.Errorf("error while finding product: %w", err)
    }
	
	 return productData, claim, nil

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