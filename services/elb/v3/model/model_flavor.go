package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flavor 负载均衡器规格信息。
type Flavor struct {

	// 规格ID。
	Id string `json:"id"`

	Info *FlavorInfo `json:"info"`

	// 规格名称。  规格名称与控制台展示的对应关系如下：  网络型有如下规格：   - L4_flavor.elb.s1.small: 小型 I   - L4_flavor.elb.s2.small: 小型 II   - L4_flavor.elb.s1.medium: 中型 I   - L4_flavor.elb.s2.medium: 中型 II   - L4_flavor.elb.s1.large: 大型 I   - L4_flavor.elb.s2.large: 大型 II  应用型有如下规格：   - L7_flavor.elb.s1.small: 小型 I   - L7_flavor.elb.s2.small: 小型 II   - L7_flavor.elb.s1.medium: 中型 I   - L7_flavor.elb.s2.medium: 中型 II   - L7_flavor.elb.s1.large: 大型 I   - L7_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.s1.extra-large: 超大型 I   - L7_flavor.elb.s2.extra-large: 超大型 II
	Name string `json:"name"`

	// 是否公共规格。  取值： - true表示公共规格，所有租户可见。 - false表示私有规格，为当前租户所有。
	Shared bool `json:"shared"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 规格类别。  取值： - L4和L7 表示四层和七层flavor。 - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max和L7_elastic_max 表示弹性扩缩容实例的上限规格。
	Type string `json:"type"`

	// [是否售罄。](tag:hws,hk,hws_eu,otc,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk_g42,dt_test,hcso_dt,mix,hk_sbc,hws_ocb,fcs,fcs_dt,dt) [是否无法购买该规格的LB](tag:ocb)  取值： - true：[已售罄，将](tag:hws,hk,hws_eu,otc,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk_g42,dt_test,hcso_dt,mix,hk_sbc,hws_ocb,fcs,fcs_dt,dt)无法购买该规格的LB。 - false：[未售罄，](tag:hws,hk,hws_eu,otc,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk_g42,dt_test,hcso_dt,mix,hk_sbc,hws_ocb,fcs,fcs_dt,dt)可购买该规格的LB。
	FlavorSoldOut bool `json:"flavor_sold_out"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
