package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfo struct {

	// 下一次读取位置
	NextMarker int64 `json:"next_marker" xml:"next_marker"`

	// 上一次读取位置
	PreviousMarker int64 `json:"previous_marker" xml:"previous_marker"`

	// 当前页总数
	CurrentCount int64 `json:"current_count" xml:"current_count"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
