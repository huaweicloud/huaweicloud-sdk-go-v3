package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TermTenantCssClusterDto 最终租户css集群详情
type TermTenantCssClusterDto struct {

	// css集群id
	CssId *string `json:"css_id,omitempty"`

	// css集群名称
	Name *string `json:"name,omitempty"`

	// 集群是否可用
	IsActive *bool `json:"is_active,omitempty"`
}

func (o TermTenantCssClusterDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TermTenantCssClusterDto struct{}"
	}

	return strings.Join([]string{"TermTenantCssClusterDto", string(data)}, " ")
}
