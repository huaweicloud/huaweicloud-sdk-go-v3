package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResultDetailResponseInfo 操作详情信息
type ResultDetailResponseInfo struct {

	// **参数解释**： 告警事件关键字，仅用于告警白名单 **取值范围**： 字符长度0-128位
	Keyword *string `json:"keyword,omitempty"`

	// **参数解释**： 告警事件hash，仅用于告警白名单 **取值范围**： 字符长度64位（SHA256哈希）
	Hash *string `json:"hash,omitempty"`
}

func (o ResultDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResultDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"ResultDetailResponseInfo", string(data)}, " ")
}
