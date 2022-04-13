package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateInstanceV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`

	Body *InstanceModReq `json:"body,omitempty"`
}

func (o UpdateInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceV2Request struct{}"
	}

	return strings.Join([]string{"UpdateInstanceV2Request", string(data)}, " ")
}
