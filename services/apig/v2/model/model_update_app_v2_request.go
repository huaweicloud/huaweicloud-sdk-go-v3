package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAppV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`

	Body *AppCreate `json:"body,omitempty"`
}

func (o UpdateAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppV2Request struct{}"
	}

	return strings.Join([]string{"UpdateAppV2Request", string(data)}, " ")
}
