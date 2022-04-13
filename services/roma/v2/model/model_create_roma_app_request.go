package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRomaAppRequest struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *CreateAppReq `json:"body,omitempty"`
}

func (o CreateRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRomaAppRequest struct{}"
	}

	return strings.Join([]string{"CreateRomaAppRequest", string(data)}, " ")
}
