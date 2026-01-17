package workflowstore

import (
	"errors"
	"goAssignment/model"
	"sync"
)

var workflows = make(map[string]*model.WorkFlow)
var mut sync.Mutex
// mut => pointer   we use it to pervent multiple gorutine action on single data

func CreateWorkFlow(wf *model.WorkFlow){ 
	mut.Lock()
	defer mut.Unlock()

	workflows[wf.WorkFlowId]=wf
}


func UpdateWorkFlow(workFlowId string, status string, productData any ) error{

	mut.Lock()
	defer mut.Unlock()

	wf, ok := workflows[workFlowId]
	if !ok{
		return  errors.New("Workflow not found")
	}

	wf.Status=status
	wf.Product=productData

	return  nil
}

func GetWorkFlow(workflowId string)(any, error){

	mut.Lock()
	defer mut.Unlock()

	wf, ok := workflows[workflowId]

	if !ok{
		return  nil, errors.New("WorkFlow not found")
	}

	return wf, nil
}
