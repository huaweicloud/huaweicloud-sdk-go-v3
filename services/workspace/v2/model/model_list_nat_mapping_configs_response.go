package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatMappingConfigsResponse Response Object
type ListNatMappingConfigsResponse struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 是否开启nat映射。
	NatOn *bool `json:"nat_on,omitempty"`

	// vag Ip列表。
	VagIps *[]string `json:"vag_ips,omitempty"`

	// NAT映射配置表。
	NatVagMaps     *[]NatMappingConfig `json:"nat_vag_maps,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListNatMappingConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatMappingConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListNatMappingConfigsResponse", string(data)}, " ")
}
