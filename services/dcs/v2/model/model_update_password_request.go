package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePasswordRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *ModifyInstancePasswordBody `json:"body,omitempty"`
}

func (o UpdatePasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordRequest struct{}"
	}

	return strings.Join([]string{"UpdatePasswordRequest", string(data)}, " ")
}
