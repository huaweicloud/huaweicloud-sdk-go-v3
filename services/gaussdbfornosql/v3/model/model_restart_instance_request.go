package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartInstanceRequest Request Object
type RestartInstanceRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *RestartInstanceRequestBody `json:"body,omitempty"`
}

func (o RestartInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestartInstanceRequest", string(data)}, " ")
}
