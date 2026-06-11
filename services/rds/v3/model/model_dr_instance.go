package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DrInstance struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 实例名称
	InstanceName string `json:"instance_name"`
}

func (o DrInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DrInstance struct{}"
	}

	return strings.Join([]string{"DrInstance", string(data)}, " ")
}
