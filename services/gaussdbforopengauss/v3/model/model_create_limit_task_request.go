package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLimitTaskRequest Request Object
type CreateLimitTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateLimitTaskRequestBody `json:"body,omitempty"`
}

func (o CreateLimitTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLimitTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateLimitTaskRequest", string(data)}, " ")
}
