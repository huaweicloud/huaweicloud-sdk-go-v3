package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Location 策略中的位置，形式分别为JSON表示的路径和相应的行列范围。
type Location struct {

	// 策略中的路径，表示为路径元素的有序序列。
	Path []PathElement `json:"path"`

	Span *Span `json:"span"`
}

func (o Location) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Location struct{}"
	}

	return strings.Join([]string{"Location", string(data)}, " ")
}
