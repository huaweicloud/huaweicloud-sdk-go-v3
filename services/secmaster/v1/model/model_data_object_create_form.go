package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectCreateForm 数据对象创建体
type DataObjectCreateForm struct {

	// 唯一标识ID
	Id *string `json:"id,omitempty"`

	// 对齐的模板版本号，默认传1
	FormatVersion *int32 `json:"format_version,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Type *string `json:"type,omitempty"`

	// 触发标志
	TriggerFlag *bool `json:"trigger_flag,omitempty"`

	// 单个数据对象json体，键值根据对应数据类所需的字段添加
	DataObject *interface{} `json:"data_object,omitempty"`
}

func (o DataObjectCreateForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectCreateForm struct{}"
	}

	return strings.Join([]string{"DataObjectCreateForm", string(data)}, " ")
}
