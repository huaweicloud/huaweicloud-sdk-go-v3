package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportIndicatorsRequestBody struct {
	DataObjectFiltersForm *DataobjectSearch `json:"data_object_filters_form,omitempty"`

	// 导出字段列表
	Title *[]string `json:"title,omitempty"`
}

func (o ExportIndicatorsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIndicatorsRequestBody struct{}"
	}

	return strings.Join([]string{"ExportIndicatorsRequestBody", string(data)}, " ")
}
