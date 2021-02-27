package model

import (
	"encoding/json"

	"strings"
)

type QuotaSimpleInfo struct {
	// 精英服务商ID。
	Id string `json:"id"`
	// 分配给精英服务商的代金券额度ID。
	QuotaId string `json:"quota_id"`
}

func (o QuotaSimpleInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaSimpleInfo struct{}"
	}

	return strings.Join([]string{"QuotaSimpleInfo", string(data)}, " ")
}
