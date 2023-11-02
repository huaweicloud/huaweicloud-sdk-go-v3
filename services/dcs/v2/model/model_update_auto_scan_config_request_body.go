package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoScanConfigRequestBody 修改自动扫描配置
type UpdateAutoScanConfigRequestBody struct {

	// 启用自动扫描
	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	// 首次扫描时间，例如：2023-07-07T15:00:05.000
	FirstScanAt *string `json:"first_scan_at,omitempty"`

	// 间隔时间(秒)
	Interval *int32 `json:"interval,omitempty"`

	// 超时时间(秒)
	Timeout *int32 `json:"timeout,omitempty"`

	// 扫描密钥计数
	ScanKeysCount *int32 `json:"scan_keys_count,omitempty"`
}

func (o UpdateAutoScanConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoScanConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAutoScanConfigRequestBody", string(data)}, " ")
}
