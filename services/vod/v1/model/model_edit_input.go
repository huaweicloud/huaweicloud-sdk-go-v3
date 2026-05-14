package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditInput struct {

	// 媒资ID
	AssetId string `json:"asset_id"`

	// 剪切开始时间，单位：秒，最大长度支持32。
	TimelineStart string `json:"timeline_start"`

	// 剪切结束时间，单位：秒，最大长度支持32。
	TimelineEnd string `json:"timeline_end"`
}

func (o EditInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditInput struct{}"
	}

	return strings.Join([]string{"EditInput", string(data)}, " ")
}
