package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 租户配额对象，不包含最大值和最小值。
type QuotaNoLimit struct {

	// 配额资源列表
	Resources []ResourceNoLimit `json:"resources"`
}

func (o QuotaNoLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaNoLimit struct{}"
	}

	return strings.Join([]string{"QuotaNoLimit", string(data)}, " ")
}
