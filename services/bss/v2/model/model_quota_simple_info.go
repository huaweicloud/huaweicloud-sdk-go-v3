package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaSimpleInfo struct {

	// 云经销商ID。
	Id string `json:"id"`

	// 分配给云经销商的代金券额度ID。
	QuotaId string `json:"quota_id"`
}

func (o QuotaSimpleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaSimpleInfo struct{}"
	}

	return strings.Join([]string{"QuotaSimpleInfo", string(data)}, " ")
}
