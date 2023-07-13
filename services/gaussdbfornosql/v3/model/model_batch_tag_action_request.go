package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchTagActionRequest Request Object
type BatchTagActionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchTagActionRequestBody `json:"body,omitempty"`
}

func (o BatchTagActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchTagActionRequest struct{}"
	}

	return strings.Join([]string{"BatchTagActionRequest", string(data)}, " ")
}
