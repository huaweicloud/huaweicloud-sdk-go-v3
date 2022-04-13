package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RestartManagerRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o RestartManagerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartManagerRequest struct{}"
	}

	return strings.Join([]string{"RestartManagerRequest", string(data)}, " ")
}
