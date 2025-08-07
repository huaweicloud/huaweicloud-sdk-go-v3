package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MarketplaceEngineProduct struct {

	// 引擎ID。
	EngineId *string `json:"engine_id,omitempty"`

	// 引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 云市场规格ID。
	SpecCode *string `json:"spec_code,omitempty"`

	// 云市场商品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 服务商名称。
	BpName *string `json:"bp_name,omitempty"`

	// 服务商ID。
	BpDomainId *string `json:"bp_domain_id,omitempty"`

	// 支持的实例类型。  - Single: 单机实例 - Ha: 主备实例 - Replica: 只读实例 - All: 以上类型都支持
	InstanceMode *string `json:"instance_mode,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 用户许可。
	UserLicenseAgreement *string `json:"user_license_agreement,omitempty"`

	// 许可详情列表。
	Agreements *[]Agreement `json:"agreements,omitempty"`
}

func (o MarketplaceEngineProduct) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarketplaceEngineProduct struct{}"
	}

	return strings.Join([]string{"MarketplaceEngineProduct", string(data)}, " ")
}
