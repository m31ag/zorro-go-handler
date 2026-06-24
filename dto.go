package zorro

type Variable struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type InputDto struct {
	Job                 string              `json:"job"`
	ServiceTaskId       string              `json:"serviceTaskId"`
	ServiceTaskKey      *string             `json:"serviceTaskKey"`
	ProcessInstanceId   string              `json:"processInstanceId"`
	ProcessDefinitionId string              `json:"processDefinitionId"`
	Variables           map[string]Variable `json:"variables"`
}

type OutputDto struct {
	ServiceTaskId string              `json:"serviceTaskId"`
	Variables     map[string]Variable `json:"variables"`
}

// proxy
type CompleteTaskRequest struct {
	Variables []Variable `json:"variables"`
}

type FailTaskRequest struct {
	Message string `json:"message"`
	Retries int    `json:"retries"`
}
