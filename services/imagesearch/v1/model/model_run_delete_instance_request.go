package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunDeleteInstanceRequest struct {
	// 实例名称。

	InstanceName string `json:"instance_name"`
}

func (o RunDeleteInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunDeleteInstanceRequest", string(data)}, " ")
}
