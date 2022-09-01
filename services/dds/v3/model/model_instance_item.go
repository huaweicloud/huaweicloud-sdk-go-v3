package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceItem struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 实例名称
	InstanceName string `json:"instance_name" xml:"instance_name"`

	// 标签列表。如果没有标签，默认为空数组。
	Tags []InstanceItemTagItem `json:"tags" xml:"tags"`
}

func (o InstanceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceItem struct{}"
	}

	return strings.Join([]string{"InstanceItem", string(data)}, " ")
}
