package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResponseDdosQuota 配额信息
type QuotaResponseDdosQuota struct {

	// 当前用户已使用配额
	Current *int32 `json:"current,omitempty"`

	// 最大配额
	Quota *int32 `json:"quota,omitempty"`
}

func (o QuotaResponseDdosQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResponseDdosQuota struct{}"
	}

	return strings.Join([]string{"QuotaResponseDdosQuota", string(data)}, " ")
}
