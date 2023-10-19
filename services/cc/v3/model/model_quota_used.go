package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaUsed struct {

	// 已使用配额。
	Used int32 `json:"used"`
}

func (o QuotaUsed) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaUsed struct{}"
	}

	return strings.Join([]string{"QuotaUsed", string(data)}, " ")
}
