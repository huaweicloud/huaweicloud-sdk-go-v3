package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckAppV2Request struct {
	// 实例ID

	InstanceId string `json:"instance_id"`
	// 应用编号

	AppId string `json:"app_id"`
}

func (o CheckAppV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAppV2Request struct{}"
	}

	return strings.Join([]string{"CheckAppV2Request", string(data)}, " ")
}
