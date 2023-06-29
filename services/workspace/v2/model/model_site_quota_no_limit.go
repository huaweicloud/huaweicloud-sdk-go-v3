package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteQuotaNoLimit 单个站点的配额
type SiteQuotaNoLimit struct {

	// 配额资源列表
	Resources []ResourceNoLimit `json:"resources"`

	// 站点ID
	SiteId *string `json:"site_id,omitempty"`
}

func (o SiteQuotaNoLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteQuotaNoLimit struct{}"
	}

	return strings.Join([]string{"SiteQuotaNoLimit", string(data)}, " ")
}
