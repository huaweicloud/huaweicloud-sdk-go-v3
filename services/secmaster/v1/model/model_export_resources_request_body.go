package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportResourcesRequestBody 导出资产列表请求body体
type ExportResourcesRequestBody struct {
	DataObjectFiltersForm *ResourceDataobjectSearch `json:"data_object_filters_form"`

	// 导出字段
	Title []string `json:"title"`
}

func (o ExportResourcesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportResourcesRequestBody struct{}"
	}

	return strings.Join([]string{"ExportResourcesRequestBody", string(data)}, " ")
}
