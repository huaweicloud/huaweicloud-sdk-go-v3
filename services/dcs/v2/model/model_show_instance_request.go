package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowInstanceRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRequest", string(data)}, " ")
}
