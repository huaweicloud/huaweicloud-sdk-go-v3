package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteVhostsRequest Request Object
type BatchDeleteVhostsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *BatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteVhostsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteVhostsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteVhostsRequest", string(data)}, " ")
}
