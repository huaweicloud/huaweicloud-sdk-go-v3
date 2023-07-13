package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreInstanceRequest Request Object
type RestoreInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RestoreInstanceBody `json:"body,omitempty"`
}

func (o RestoreInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreInstanceRequest", string(data)}, " ")
}
