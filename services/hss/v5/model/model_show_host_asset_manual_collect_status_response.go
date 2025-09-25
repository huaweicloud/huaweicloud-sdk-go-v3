package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostAssetManualCollectStatusResponse Response Object
type ShowHostAssetManualCollectStatusResponse struct {

	// **参数解释**： 手动检测状态 **取值范围**： - neverscan：从未扫描过 - scanning：扫描中 - scanned：扫描完成 - failed：扫描失败 - longscanning：扫描超时
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**： 检测完成时间 **取值范围**： 最小值0，最大值2^63-1
	ScannedTime    *int64 `json:"scanned_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHostAssetManualCollectStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostAssetManualCollectStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowHostAssetManualCollectStatusResponse", string(data)}, " ")
}
