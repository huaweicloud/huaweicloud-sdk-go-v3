package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScanStatusResponse Response Object
type ShowScanStatusResponse struct {

	// **参数解释**： 手动检测状态 **取值范围**： - neverscan：无扫描任务 - scanning：扫描中 - scaned：扫描完成 - failed：扫描失败 - over_time：扫描超时（扫描时间超过30分钟） - longscanning：扫描超时（扫描时间超过60分钟）
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**： 检测完成时间 **取值范围**： 不涉及
	ScanedTime     *int64 `json:"scaned_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowScanStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScanStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowScanStatusResponse", string(data)}, " ")
}
