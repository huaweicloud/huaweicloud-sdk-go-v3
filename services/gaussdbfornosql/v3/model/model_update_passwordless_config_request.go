package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePasswordlessConfigRequest Request Object
type UpdatePasswordlessConfigRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdatePasswordlessConfigRequestBody `json:"body,omitempty"`
}

func (o UpdatePasswordlessConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePasswordlessConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdatePasswordlessConfigRequest", string(data)}, " ")
}
