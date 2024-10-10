package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor 负载均衡器规格信息。
type Flavor struct {

	// 参数解释：规格ID。
	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`

	// 参数解释：规格名称。  取值范围：  网络型有如下规格：   - L4_flavor.elb.s1.small: 小型 I   - L4_flavor.elb.s2.small: 小型 II   - L4_flavor.elb.s1.medium: 中型 I   - L4_flavor.elb.s2.medium: 中型 II   - L4_flavor.elb.s1.large: 大型 I   - L4_flavor.elb.s2.large: 大型 II   - L4_flavor.elb.pro.max: 四层弹性规格  应用型有如下规格：   - L7_flavor.elb.s1.small: 小型 I   - L7_flavor.elb.s2.small: 小型 II   - L7_flavor.elb.s1.medium: 中型 I   - L7_flavor.elb.s2.medium: 中型 II   - L7_flavor.elb.s1.large: 大型 I   - L7_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.s1.extra-large: 超大型 I   - L7_flavor.elb.s2.extra-large: 超大型 II   - L7_flavor.elb.pro.max: 七层弹性规格
	Name string `json:"name"`

	// 参数解释：是否公共规格。  取值范围： - true表示公共规格，所有租户可见。 - false表示私有规格，为当前租户所有。
	Shared bool `json:"shared"`

	// 参数解释：项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：规格类别。  取值范围：   - L4和L7 表示四层网络型和七层应用型flavor。   [- gateway 表示网关型LB的flavor，目前只支持弹性计费类型。当前仅支持欧洲局点。](tag:hws_eu)   - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。已废弃，请勿使用。   - L4_elastic_max、L7_elastic_max[和gateway_elastic_max](tag:hws_eu) 表示弹性扩缩容实例的上限规格。
	Type string `json:"type"`

	// 参数解释：[是否售罄。](tag:hws,hk,hws_eu,otc,tlf,sbc,g42,hk_g42,dt_test,mix,hk_sbc,hws_ocb,dt) [是否无法购买该规格的LB](tag:ctc,cmcc,ocb,tm,fcs,fcs_dt,hcso,hcso_dt,hk_vdf)  取值范围： - true：[已售罄，将](tag:hws,hk,hws_eu,otc,tlf,sbc,g42,hk_g42,dt_test,mix,hk_sbc,hws_ocb,dt)无法购买该规格的LB。 - false：[未售罄，](tag:hws,hk,hws_eu,otc,tlf,sbc,g42,hk_g42,dt_test,mix,hk_sbc,hws_ocb,dt)可购买该规格的LB。
	FlavorSoldOut bool `json:"flavor_sold_out"`

	// 参数解释：可用区组，如：center
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 参数解释：可用区组编码。  取值范围：0表示center，21表示homezone。
	Category *int32 `json:"category,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
