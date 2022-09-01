package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObsScanResult struct {

	// 扫描结果总数
	Total *int32 `json:"total,omitempty" xml:"total"`

	// OBS扫描结果列表
	DbScanResults *[]ObsScanResultInfo `json:"db_scan_results,omitempty" xml:"db_scan_results"`
}

func (o ObsScanResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsScanResult struct{}"
	}

	return strings.Join([]string{"ObsScanResult", string(data)}, " ")
}
