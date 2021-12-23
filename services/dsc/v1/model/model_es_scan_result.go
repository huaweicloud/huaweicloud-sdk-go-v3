package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EsScanResult struct {
	// 扫描结果总数

	Total *int32 `json:"total,omitempty"`
	// ES扫描结果列表

	DbScanResults *[]EsScanResultInfo `json:"db_scan_results,omitempty"`
}

func (o EsScanResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EsScanResult struct{}"
	}

	return strings.Join([]string{"EsScanResult", string(data)}, " ")
}
