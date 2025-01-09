package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatMappingConfigsReq 修改租户NAT映射配置请求体。
type UpdateNatMappingConfigsReq struct {

	// 站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 是否开启nat映射。
	NatOn *bool `json:"nat_on,omitempty"`

	// NAT映射配置表。
	NatVagMaps *[]UpdateNatMappingConfig `json:"nat_vag_maps,omitempty"`
}

func (o UpdateNatMappingConfigsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatMappingConfigsReq struct{}"
	}

	return strings.Join([]string{"UpdateNatMappingConfigsReq", string(data)}, " ")
}
