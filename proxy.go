package zorro

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Proxy interface {
	CompleteTask(dto OutputDto) error
	FailTask(id string, message string) error
}
type proxy struct {
	client *resty.Client
	url    string
}

func NewProxy(url string) Proxy {
	return &proxy{
		client: resty.New(),
		url:    url,
	}
}

func (p *proxy) CompleteTask(dto OutputDto) error {
	vars := make([]Variable, 0, len(dto.Variables))
	for _, v := range dto.Variables {
		vars = append(vars, v)
	}

	req := CompleteTaskRequest{Variables: vars}
	resp, err := p.client.R().
		SetBody(req).
		Post(fmt.Sprintf("%s/service-tasks/%s/complete", p.url, dto.ServiceTaskId))
	if err != nil {
		return fmt.Errorf("complete task request: %w", err)
	}
	if resp == nil {
		return errors.New("response is nil")
	}
	return nil
}

func (p *proxy) FailTask(id string, message string) error {
	req := FailTaskRequest{Message: message, Retries: 0}
	resp, err := p.client.R().
		SetBody(req).
		Post(fmt.Sprintf("%s/service-tasks/%s/fail", p.url, id))
	if err != nil {
		return fmt.Errorf("fail task request: %w", err)
	}
	if resp == nil {
		return errors.New("response is nil")
	}
	return nil
}
