package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor **参数解释**：负载均衡器规格信息。
type Flavor struct {

	// **参数解释**：规格ID。  **取值范围**：不涉及
	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`

	// **参数解释**：规格名称。  **取值范围**： [网络型有如下规格：   - L4_flavor.elb.s1.small: 小型 I   - L4_flavor.elb.s2.small: 小型 II   - L4_flavor.elb.s1.medium: 中型 I   - L4_flavor.elb.s2.medium: 中型 II   - L4_flavor.elb.s1.large: 大型 I   - L4_flavor.elb.s2.large: 大型 II   - L4_flavor.elb.pro.max: 四层弹性规格  应用型有如下规格：   - L7_flavor.elb.s1.small: 小型 I   - L7_flavor.elb.s2.small: 小型 II   - L7_flavor.elb.s1.medium: 中型 I   - L7_flavor.elb.s2.medium: 中型 II   - L7_flavor.elb.s1.large: 大型 I   - L7_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.s1.extra-large: 超大型 I   - L7_flavor.elb.s2.extra-large: 超大型 II   - L7_flavor.elb.pro.max: 七层弹性规格](tag:hws,hws_hk,hws_eu,hws_eu_wb,hws_test,fcs,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg,g42,hk_g42)  [网络型有如下规格：   - L4_flavor.elb.s1.small: 小型 I   - L4_flavor.elb.s2.small: 小型 II   - L4_flavor.elb.s1.medium: 中型 I   - L4_flavor.elb.s2.medium: 中型 II   - L4_flavor.elb.s1.large: 大型 I   - L4_flavor.elb.s2.large: 大型 II   - L4_flavor.elb.pro.max: 四层弹性规格  应用型有如下规格：   - L7_flavor.elb.s1.small: 小型 I   - L7_flavor.elb.s2.small: 小型 II   - L7_flavor.elb.s1.medium: 中型 I   - L7_flavor.elb.s2.medium: 中型 II   - L7_flavor.elb.s1.large: 大型 I   - L7_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.pro.s1.small：弹性规格小型I。已不支持，请勿使用。   - L7_flavor.elb.pro.s1.medium：弹性规格中型I。已不支持，请勿使用。   - L7_flavor.elb.pro.s1.large：弹性规格大型I。已不支持，请勿使用。   - L7_flavor.elb.pro.max: 七层弹性规格](tag:dt,hcso_dt)
	Name string `json:"name"`

	// **参数解释**：是否公共规格。  **取值范围**： - true表示公共规格，所有租户可见。 - false表示私有规格，为当前租户所有。
	Shared bool `json:"shared"`

	// **参数解释**：项目ID。获取方式请参见[获取项目ID](elb_fl_0008.xml)。  **取值范围**：长度为32个字符，由小写字母和数字组成。
	ProjectId string `json:"project_id"`

	// **参数解释**：规格类别。  **取值范围**：   - L4和L7 表示四层网络型和七层应用型flavor。   - gwlb 表示网关型LB的flavor。   - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。已废弃，请勿使用。   - L4_elastic_max、L7_elastic_max和gwlb_elastic_max表示弹性扩缩容实例的上限规格。
	Type string `json:"type"`

	// **参数解释**：[是否售罄。](tag:hws)[是否无法购买该规格的LB。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,fcs,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg,ct)  **取值范围**： - true：[已售罄，将](tag:hws)无法购买该规格的LB。 - false：[未售罄，](tag:hws)可购买该规格的LB。
	FlavorSoldOut bool `json:"flavor_sold_out"`

	// **参数解释**：公网边界组。  **取值范围**： - center：表示中心站点的公网边界组 - 边缘站点名称：表示边缘站点的公网边界组
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// **参数解释**：可用区子类型编码。该字段主要用于区分在边缘场景下，边缘AZ的类型。  **取值范围**：0表示center，21表示homezone，41表示IES。
	Category *int32 `json:"category,omitempty"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
