package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateExpireAutoScanConfigResponse Response Object
type UpdateExpireAutoScanConfigResponse struct {

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 是否开启自动扫描
	EnableAutoScan *bool `json:"enable_auto_scan,omitempty"`

	// 首次扫描时间，例如：2023-07-07T15:00:05.000
	FirstScanAt *string `json:"first_scan_at,omitempty"`

	// 间隔时间(秒)
	Interval *int32 `json:"interval,omitempty"`

	// 超时时间(秒)
	Timeout *int32 `json:"timeout,omitempty"`

	// 扫描数
	ScanKeysCount *int32 `json:"scan_keys_count,omitempty"`

	// 更新时间，例如：2023-06-15T06:20:13.283Z
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateExpireAutoScanConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateExpireAutoScanConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateExpireAutoScanConfigResponse", string(data)}, " ")
}
