package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataObjectDeleteForm 数据对象删除体
type DataObjectDeleteForm struct {

	// 批量删除ID列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o DataObjectDeleteForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectDeleteForm struct{}"
	}

	return strings.Join([]string{"DataObjectDeleteForm", string(data)}, " ")
}
