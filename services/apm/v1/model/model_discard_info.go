package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiscardInfo 丢弃的信息。
type DiscardInfo struct {

	// 类型。
	Type *string `json:"type,omitempty"`

	// 数量。
	Count *int32 `json:"count,omitempty"`

	// 总时间。
	TotalTime *int64 `json:"totalTime,omitempty"`
}

func (o DiscardInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscardInfo struct{}"
	}

	return strings.Join([]string{"DiscardInfo", string(data)}, " ")
}
