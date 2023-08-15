package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsScanResult OBS扫描结果
type ObsScanResult struct {

	// 扫描结果总数
	Total *int32 `json:"total,omitempty"`

	// OBS扫描结果列表
	DbScanResults *[]ObsScanResultInfo `json:"db_scan_results,omitempty"`
}

func (o ObsScanResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsScanResult struct{}"
	}

	return strings.Join([]string{"ObsScanResult", string(data)}, " ")
}
