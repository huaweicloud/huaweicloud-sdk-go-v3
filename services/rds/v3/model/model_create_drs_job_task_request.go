package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDrsJobTaskRequest Request Object
type CreateDrsJobTaskRequest struct {

	// 实例id，源实例
	InstanceId string `json:"instance_id"`

	Body *CreateDrsTaskReq `json:"body,omitempty"`
}

func (o CreateDrsJobTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrsJobTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateDrsJobTaskRequest", string(data)}, " ")
}
