package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultDetailResponseInfo 操作详情信息
type ResultDetailResponseInfo struct {

	// 告警事件关键字，仅用于告警白名单
	Keyword *string `json:"keyword,omitempty"`

	// 告警事件hash，仅用于告警白名单
	Hash *string `json:"hash,omitempty"`
}

func (o ResultDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"ResultDetailResponseInfo", string(data)}, " ")
}
