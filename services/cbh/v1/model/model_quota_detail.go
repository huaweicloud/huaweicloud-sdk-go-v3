package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 当前租户CBH的配额信息
type QuotaDetail struct {
	// 中文配额描述

	ZhCn *string `json:"zh_cn,omitempty"`
	// 英文配额描述

	EnUs *string `json:"en_us,omitempty"`
	// 租户剩余配额数量

	Remaining int32 `json:"remaining"`
}

func (o QuotaDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetail struct{}"
	}

	return strings.Join([]string{"QuotaDetail", string(data)}, " ")
}
