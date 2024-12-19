package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeVolumeRequestBody 退订包周期云硬盘的请求body体
type UnsubscribeVolumeRequestBody struct {

	// 退订包周期云硬盘的结果。
	VolumeIds []string `json:"volume_ids"`
}

func (o UnsubscribeVolumeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"UnsubscribeVolumeRequestBody", string(data)}, " ")
}
