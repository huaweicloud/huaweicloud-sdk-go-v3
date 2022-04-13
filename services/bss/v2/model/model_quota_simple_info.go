package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaSimpleInfo struct {
	// 精英服务商ID。

	Id string `json:"id"`
	// 分配给精英服务商的代金券额度ID。

	QuotaId string `json:"quota_id"`
}

func (o QuotaSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaSimpleInfo struct{}"
	}

	return strings.Join([]string{"QuotaSimpleInfo", string(data)}, " ")
}
