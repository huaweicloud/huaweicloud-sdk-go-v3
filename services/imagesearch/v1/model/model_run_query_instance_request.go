package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunQueryInstanceRequest struct {

	// 实例名称。
	InstanceName string `json:"instance_name" xml:"instance_name"`
}

func (o RunQueryInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryInstanceRequest struct{}"
	}

	return strings.Join([]string{"RunQueryInstanceRequest", string(data)}, " ")
}
