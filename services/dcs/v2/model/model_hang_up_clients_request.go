package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HangUpClientsRequest Request Object
type HangUpClientsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *HangUpClientsRequestBody `json:"body,omitempty"`
}

func (o HangUpClientsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HangUpClientsRequest struct{}"
	}

	return strings.Join([]string{"HangUpClientsRequest", string(data)}, " ")
}
