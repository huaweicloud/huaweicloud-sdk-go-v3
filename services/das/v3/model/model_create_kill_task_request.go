package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKillTaskRequest Request Object
type CreateKillTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateKillTaskRequestBody `json:"body,omitempty"`
}

func (o CreateKillTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKillTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateKillTaskRequest", string(data)}, " ")
}
