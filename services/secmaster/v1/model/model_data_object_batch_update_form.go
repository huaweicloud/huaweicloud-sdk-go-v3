package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectBatchUpdateForm 数据对象批量更新对象
type DataObjectBatchUpdateForm struct {

	// 批量更新ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`

	// 触发标志
	TriggerFlag *bool `json:"trigger_flag,omitempty"`

	// 数据对象
	DataObject *interface{} `json:"data_object,omitempty"`
}

func (o DataObjectBatchUpdateForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectBatchUpdateForm struct{}"
	}

	return strings.Join([]string{"DataObjectBatchUpdateForm", string(data)}, " ")
}
