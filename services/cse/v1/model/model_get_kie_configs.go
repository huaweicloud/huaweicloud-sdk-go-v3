package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetKieConfigs struct {

	// 配置项的id。
	Id *string `json:"id,omitempty"`

	// 配置项的key。
	Key *string `json:"key,omitempty"`

	// 配置项的标签。
	Labels *interface{} `json:"labels,omitempty"`

	// 配置项的值。
	Value *string `json:"value,omitempty"`

	// 配置项value的类型。
	ValueType *string `json:"value_type,omitempty"`

	// 配置项的状态。
	Status *string `json:"status,omitempty"`

	// 创建时间。
	CreateTime *int32 `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *int32 `json:"update_time,omitempty"`

	// 创建配置的版本号
	CreateRevision *int64 `json:"create_revision,omitempty"`

	// 修改配置的版本号
	UpdateRevision *int64 `json:"update_revision,omitempty"`
}

func (o GetKieConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetKieConfigs struct{}"
	}

	return strings.Join([]string{"GetKieConfigs", string(data)}, " ")
}
