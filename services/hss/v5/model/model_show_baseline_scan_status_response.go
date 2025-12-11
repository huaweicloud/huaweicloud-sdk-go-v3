package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineScanStatusResponse Response Object
type ShowBaselineScanStatusResponse struct {

	// **参数解释**: 扫描状态 **约束限制**: 不涉及 **取值范围**: - neverscan: 未扫描 - scanning: 扫描中 - scanned: 扫描完成 - failed: 扫描失败  **默认取值**: 不涉及
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 扫描结束时间，单位毫秒 **约束限制**: 不涉及 **取值范围**: 大小0-4070880000000 **默认取值**: 不涉及
	ScannedTime    *int64 `json:"scanned_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowBaselineScanStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineScanStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineScanStatusResponse", string(data)}, " ")
}
