package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCommonTaskRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateCommonTaskRequestBody `json:"body,omitempty"`
}

func (o CreateCommonTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommonTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateCommonTaskRequest", string(data)}, " ")
}
