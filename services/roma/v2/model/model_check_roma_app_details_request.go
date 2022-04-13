package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckRomaAppDetailsRequest struct {
	// 应用ID

	AppId string `json:"app_id"`
	// 实例ID

	InstanceId string `json:"instance_id"`
}

func (o CheckRomaAppDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRomaAppDetailsRequest struct{}"
	}

	return strings.Join([]string{"CheckRomaAppDetailsRequest", string(data)}, " ")
}
