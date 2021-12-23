package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceResult struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 实例名称。

	InstanceName string `json:"instance_name"`
	// 标签列表。如果没有标签，默认为空数组。

	Tags []InstanceTagResult `json:"tags"`
}

func (o InstanceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceResult struct{}"
	}

	return strings.Join([]string{"InstanceResult", string(data)}, " ")
}
