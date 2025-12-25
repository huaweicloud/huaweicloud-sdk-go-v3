package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportVulnerabilitiesRequestBody 导出漏洞列表请求body体
type ExportVulnerabilitiesRequestBody struct {
	DataObjectFiltersForm *VulnerabilityDataObjectSearch `json:"data_object_filters_form,omitempty"`

	// 导出字段
	Title *[]string `json:"title,omitempty"`
}

func (o ExportVulnerabilitiesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportVulnerabilitiesRequestBody struct{}"
	}

	return strings.Join([]string{"ExportVulnerabilitiesRequestBody", string(data)}, " ")
}
