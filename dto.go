package zorro

import (
	"fmt"
)

type Variable struct {
	Name  string       `json:"name"`
	Type  VariableType `json:"type"`
	Value string       `json:"value"`
}

type InputDto struct {
	Job                 string              `json:"job"`
	ServiceTaskId       string              `json:"serviceTaskId"`
	ServiceTaskKey      *string             `json:"serviceTaskKey"`
	ProcessInstanceId   string              `json:"processInstanceId"`
	ProcessDefinitionId string              `json:"processDefinitionId"`
	Variables           map[string]Variable `json:"variables"`
}

// Get
// return variable by name
func (i *InputDto) Get(n string) string {
	return i.Variables[n].Value
}

// GetNotEmpty
// returns variable by name if it exists or is not empty else returns error
func (i *InputDto) GetNotEmpty(n string) (string, error) {
	v, ok := i.Variables[n]
	if !ok {
		return "", fmt.Errorf("no variable with name %s exists", n)
	}
	if v.Value == "" {
		return "", fmt.Errorf("variable with name %s is empty", n)
	}
	return v.Value, nil
}

func (i *InputDto) Vars() []Variable {
	items := make([]Variable, 0, len(i.Variables))

	for _, v := range i.Variables {
		items = append(items, v)
	}
	return items
}

type OutputDto struct {
	ServiceTaskId string     `json:"serviceTaskId"`
	Variables     []Variable `json:"variables"`
}

// proxy
type CompleteTaskRequest struct {
	Variables []Variable `json:"variables"`
}

type FailTaskRequest struct {
	Message string `json:"message"`
	Retries int    `json:"retries"`
}
