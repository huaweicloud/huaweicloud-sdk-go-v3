package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceNameRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *UpdateInstanceNameRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateInstanceNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequest", string(data)}, " ")
}
