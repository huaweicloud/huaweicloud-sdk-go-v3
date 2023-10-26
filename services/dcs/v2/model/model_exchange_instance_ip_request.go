package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExchangeInstanceIpRequest Request Object
type ExchangeInstanceIpRequest struct {

	// 任务ID
	TaskId string `json:"task_id"`

	Body *IpExchangeRequest `json:"body,omitempty"`
}

func (o ExchangeInstanceIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExchangeInstanceIpRequest struct{}"
	}

	return strings.Join([]string{"ExchangeInstanceIpRequest", string(data)}, " ")
}
