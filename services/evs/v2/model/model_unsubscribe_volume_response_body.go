package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeVolumeResponseBody 退订包周期云硬盘的响应body体
type UnsubscribeVolumeResponseBody struct {

	// 退订包周期云硬盘的结果。
	Results []UnsubscribeVolume `json:"results"`
}

func (o UnsubscribeVolumeResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeVolumeResponseBody struct{}"
	}

	return strings.Join([]string{"UnsubscribeVolumeResponseBody", string(data)}, " ")
}
