package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetKieConfigs struct {

	// 配置项的id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 配置项的key。
	Key *string `json:"key,omitempty" xml:"key"`

	// 配置项的标签。
	Labels *interface{} `json:"labels,omitempty" xml:"labels"`

	// 配置项的值。
	Value *string `json:"value,omitempty" xml:"value"`

	// 配置项value的类型。
	ValueType *string `json:"value_type,omitempty" xml:"value_type"`

	// 配置项的状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 创建时间。
	CreateTime *int32 `json:"create_time,omitempty" xml:"create_time"`

	// 更新时间。
	UpdateTime *int32 `json:"update_time,omitempty" xml:"update_time"`
}

func (o GetKieConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKieConfigs struct{}"
	}

	return strings.Join([]string{"GetKieConfigs", string(data)}, " ")
}
