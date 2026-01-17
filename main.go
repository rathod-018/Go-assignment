package main

import (
	"encoding/json"
	"fmt"
	"goAssignment/model"
	"goAssignment/process"
	workflowstore "goAssignment/workFlowStore"
	"net/http"

	"github.com/google/uuid"
)


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server")
	})

	http.HandleFunc("/claims/identify", func(w http.ResponseWriter, r *http.Request) {

		// allow only post req
		if r.Method != http.MethodPost{
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req ProductReq
		if	err := json.NewDecoder(r.Body).Decode(&req); err != nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error":"invalid productId",
			})
			return
		}

		productData, claim, err := process.LoadData(req.ProductId)

		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error":err.Error(),
			})
		}

		workFlow := &model.WorkFlow{
        WorkFlowId: "wf-" + uuid.New().String(),
        Status:     "IN_PROGRESS",
    	}

		workflowstore.CreateWorkFlow(workFlow)

		go process.DetectClaims(productData, claim, workFlow.WorkFlowId)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(workFlow)
	})

	http.HandleFunc("/claims/status/{workflowId}", func(w http.ResponseWriter, r *http.Request) {
		workFlowId := r.PathValue("workflowId")
		
		workFlow, err := workflowstore.GetWorkFlow(workFlowId)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]any{
				"error":err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(workFlow)
	})


	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}


type ProductReq struct{
	ProductId 	string `json:"productId"`
}

