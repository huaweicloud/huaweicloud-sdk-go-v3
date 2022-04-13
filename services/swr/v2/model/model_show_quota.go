package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowQuota struct {
	// 配额限制

	QuotaLimit int32 `json:"quota_limit"`
	// 配额类型

	QuotaKey string `json:"quota_key"`
	// 单位

	Unit string `json:"unit"`
	// 已用配额

	Used int32 `json:"used"`
}

func (o ShowQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuota struct{}"
	}

	return strings.Join([]string{"ShowQuota", string(data)}, " ")
}
