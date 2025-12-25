package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommonDataObjectExportRequest 导出数据对象列表请求body体
type CommonDataObjectExportRequest struct {
	DataObjectFiltersForm *DataobjectSearch `json:"data_object_filters_form,omitempty"`

	// 导出动作，exportDataObject导出数据/exportTemplate导出模板
	Action *string `json:"action,omitempty"`

	// 导出字段列表
	Title *[]string `json:"title,omitempty"`
}

func (o CommonDataObjectExportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommonDataObjectExportRequest struct{}"
	}

	return strings.Join([]string{"CommonDataObjectExportRequest", string(data)}, " ")
}
