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

	// 规格名称。  规格名称与控制台展示的对应关系如下：   - L4_flavor.elb.s1.small: 小型 I   - L4_flavor.elb.s2.small: 小型 II   - L4_flavor.elb.s1.medium: 中型 I   - L4_flavor.elb.s2.medium: 中型 II   - L4_flavor.elb.s1.large: 大型 I   - L4_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.s1.small: 小型 I   - L7_flavor.elb.s2.small: 小型 II   - L7_flavor.elb.s1.medium: 中型 I   - L7_flavor.elb.s2.medium: 中型 II   - L7_flavor.elb.s1.large: 大型 I   - L7_flavor.elb.s2.large: 大型 II   - L7_flavor.elb.s1.extra-large: 超大型 I   - L7_flavor.elb.s2.extra-large: 超大型 II
	Name string `json:"name"`

	// 是否公共规格。  取值： - true表示公共规格，所有租户可见。 - false表示私有规格，为当前租户所有。
	Shared bool `json:"shared"`

	// 项目ID。
	ProjectId string `json:"project_id"`

	// 规格类别。  取值： - L4和L7 表示四层和七层flavor。 [- L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max和L7_elastic_max 表示弹性扩缩容实例的上限规格。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,fcs,hcso_dt)
	Type string `json:"type"`

	// [是否售罄。  取值： - true：已售罄，将无法购买该规格的LB。 - false：未售罄，可购买该规格的LB。](tag:hws)  [是否可用。  取值： - true：不可用，将无法创建该规格的LB。 - false：可用，可创建该规格的LB。](tag:hws_ocb,ocb,fcs,hws_hk,ctc,hcs,cmcc,tm,g42,hk_g42,hws_eu,hcso_dt,dt,dt_test)
	FlavorSoldOut bool `json:"flavor_sold_out"`
}

func (o Flavor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flavor struct{}"
	}

	return strings.Join([]string{"Flavor", string(data)}, " ")
}
