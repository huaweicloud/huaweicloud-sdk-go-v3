package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportIncidentsRequestBody 导出事件列表请求body体
type ExportIncidentsRequestBody struct {
	DataObjectFiltersForm *DataobjectSearch `json:"data_object_filters_form,omitempty"`

	// 导出字段列表
	Title *[]string `json:"title,omitempty"`
}

func (o ExportIncidentsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportIncidentsRequestBody struct{}"
	}

	return strings.Join([]string{"ExportIncidentsRequestBody", string(data)}, " ")
}
