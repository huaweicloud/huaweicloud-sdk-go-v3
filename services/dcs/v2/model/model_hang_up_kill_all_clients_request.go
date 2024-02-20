package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HangUpKillAllClientsRequest Request Object
type HangUpKillAllClientsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *HangUpKillAllClientsRequestBody `json:"body,omitempty"`
}

func (o HangUpKillAllClientsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpKillAllClientsRequest struct{}"
	}

	return strings.Join([]string{"HangUpKillAllClientsRequest", string(data)}, " ")
}
