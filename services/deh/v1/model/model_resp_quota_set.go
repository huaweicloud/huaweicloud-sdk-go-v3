package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 专属主机的配额。
type RespQuotaSet struct {
	// 配额类别。

	Resource string `json:"resource"`
	// 配额最大限制。  “-1”表示资源配额不受限制。

	HardLimit int32 `json:"hard_limit"`
	// 已使用配额数量。

	Used int32 `json:"used"`
}

func (o RespQuotaSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespQuotaSet struct{}"
	}

	return strings.Join([]string{"RespQuotaSet", string(data)}, " ")
}
