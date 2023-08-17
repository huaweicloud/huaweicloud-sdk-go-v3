package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateParam struct {

	// 数据的服务实例级唯一标识，字符长度范围为[1, 256]。
	ItemId string `json:"item_id"`

	// 数据的描述信息，字符长度范围为[1, 2048]。
	Desc *string `json:"desc,omitempty"`

	// 数据的自定义字符标签，用于进行条件过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 类型为字符串，字符长度范围为[1, 64]。
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	// 数据的自定义数值标签，用于进行条件过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_num_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 类型为数值，格式为double。
	CustomNumTags map[string]float64 `json:"custom_num_tags,omitempty"`
}

func (o UpdateParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParam struct{}"
	}

	return strings.Join([]string{"UpdateParam", string(data)}, " ")
}
