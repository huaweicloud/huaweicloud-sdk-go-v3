package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResourcesRequestBodyIacScanInfo IAC安全扫描任务特有的查询条件，任务类型为“iac_scan”可用
type ListTaskResourcesRequestBodyIacScanInfo struct {

	// 文件的扫描状态，包含如下：   - scanning：扫描中   - success：扫描成功   - failed：扫描失败
	ScanStatus *string `json:"scan_status,omitempty"`

	// 文件名称
	FileName *string `json:"file_name,omitempty"`

	// chart包名称
	ChartName *string `json:"chart_name,omitempty"`
}

func (o ListTaskResourcesRequestBodyIacScanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResourcesRequestBodyIacScanInfo struct{}"
	}

	return strings.Join([]string{"ListTaskResourcesRequestBodyIacScanInfo", string(data)}, " ")
}
