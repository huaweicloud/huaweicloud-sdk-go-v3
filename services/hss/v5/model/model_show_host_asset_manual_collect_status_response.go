package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostAssetManualCollectStatusResponse Response Object
type ShowHostAssetManualCollectStatusResponse struct {

	// 手动检测状态
	ScanStatus *string `json:"scan_status,omitempty"`

	// 检测完成时间
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
