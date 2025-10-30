package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Server 服务器详情
type Server struct {

	// 资源ID
	Id *string `json:"id,omitempty"`

	// 账号ID
	DomainId *string `json:"domain_id,omitempty"`

	// 站点ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	Status *ServerStatus `json:"status,omitempty"`

	// 商品ID
	OfferingId *string `json:"offering_id,omitempty"`

	Spec *ServerResourceSpec `json:"spec,omitempty"`

	MarketOptions *MarketOptions `json:"market_options,omitempty"`

	Location *LayoutLocation `json:"location,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o Server) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Server struct{}"
	}

	return strings.Join([]string{"Server", string(data)}, " ")
}
