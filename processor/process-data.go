package processor

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

 var claims = []map[string]interface{}{}

func ProcessData() any {
	dataBytes, dataErr :=os.ReadFile("data/data.json")
	claimBytes, climErr := os.ReadFile("data/claim.json")
	if dataErr != nil || climErr != nil{
		fmt.Println("Error while reading file")
	}
	var data map[string]any
	var claim map[string]any

	if err := json.Unmarshal(dataBytes, &data); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(claimBytes, &claim); err != nil {
		panic(err)
	}

	
 
	for key, value := range claim{
		// fmt.Printf("type of value %T\n",value)

		parts := strings.Split(value.(string),", ")

		for _,value :=range parts{
			findClaim(key, value, data)
		}
	}

	// fmt.Println(claims)

	// jsonData, err :=json.Marshal(claims)

	// if err != nil{
	// 	panic(err)
	// }

	// return string(jsonData)
	 return claims

}


func findClaim(key string, claim string, data any) {

	if m, ok := data.(map[string]interface{}); ok{
		for _, value := range m{
			findClaim(key, claim, value)
		}
		return
	}

	if s, ok := data.([]interface{}); ok{
		for _, item := range s{
			findClaim(key, claim,item)
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

			result := map[string]interface{}{
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
	ID         string 
	ClaimType  string 
	ClaimValue string 
	Status     string 
}
