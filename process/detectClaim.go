package process

import (
	workflowstore "goAssignment/workFlowStore"
	"strings"

	"github.com/google/uuid"
)


func DetectClaims(productData any, claim any, workFlowId string) {
	 claims := []map[string]any{}

	claimMap, _ := claim.(map[string]string) 
	for key, value := range claimMap{

		splitValues := strings.Split(value,", ")

		for _,value :=range splitValues{
			findClaims(key, value, productData, &claims)
		}

	}
	
	workflowstore.UpdateWorkFlow(workFlowId, "COMPLETED", productData, claims)

}

func findClaims(key string, claim string, data any, claims *[]map[string]any){
if m, ok := data.(map[string]any); ok {
		for _, value := range m {
			findClaims(key, claim, value, claims)
		}
		return
	}

	if s, ok := data.([]any); ok {
		for _, item := range s {
			findClaims(key, claim, item, claims)
		}
		return
	}

	if _, ok := data.(string); ok {
		isContains := strings.Contains(strings.ToLower(data.(string)), strings.ToLower(claim))
		if isContains {
			id := uuid.New().String()

			newClaim := Claim{
				ID:         id,
				ClaimType:  key,
				ClaimValue: claim,
				Status:     "IDENTIFIED",
			}

			result := map[string]any{
				"id":         newClaim.ID,
				"claimType":  newClaim.ClaimType,
				"claimValue": newClaim.ClaimValue,
				"status":     newClaim.Status,
			}

			*claims = append(*claims, result)
			return
		}
	}
}

type Claim struct {
	ID         string `json:"id"`
	ClaimType  string `json:"claimType"`
	ClaimValue string `json:"claimValue"`
	Status     string `json:"status"`
}