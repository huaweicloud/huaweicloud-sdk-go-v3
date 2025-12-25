package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectBatchCreateForm 数据对象批量创建对象
type DataObjectBatchCreateForm struct {

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

	// 数据对象列表
	DataObjectList *[]interface{} `json:"data_object_list,omitempty"`
}

func (o DataObjectBatchCreateForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectBatchCreateForm struct{}"
	}

	return strings.Join([]string{"DataObjectBatchCreateForm", string(data)}, " ")
}
