package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportAlertsRequestBody 导出告警列表请求body体
type ExportAlertsRequestBody struct {
	DataObjectFiltersForm *DataobjectSearch `json:"data_object_filters_form,omitempty"`

	// 导出字段列表
	Title *[]string `json:"title,omitempty"`
}

func (o ExportAlertsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportAlertsRequestBody struct{}"
	}

	return strings.Join([]string{"ExportAlertsRequestBody", string(data)}, " ")
}
