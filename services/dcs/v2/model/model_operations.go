package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Operations 操作详情
type Operations struct {

	// 操作信息
	Operation *string `json:"operation,omitempty"`

	// 是否支持该操作
	IsSupport *bool `json:"is_support,omitempty"`

	// 不支持该操作的原因ID，仅在is_support为false时返回
	CauseId *string `json:"cause_id,omitempty"`
}

func (o Operations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Operations struct{}"
	}

	return strings.Join([]string{"Operations", string(data)}, " ")
}
