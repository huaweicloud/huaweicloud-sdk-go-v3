package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuotaLimit struct {

	// 配额大小。
	QuotaLimit int32 `json:"quota_limit"`
}

func (o QuotaLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaLimit struct{}"
	}

	return strings.Join([]string{"QuotaLimit", string(data)}, " ")
}
