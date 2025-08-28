package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// **参数解释**：规格ID。 支持多值查询，查询条件格式：*id=xxx&id=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Id *[]string `json:"id,omitempty"`

	// **参数解释**：规格名称。 支持多值查询，查询条件格式：*name=xxx&name=xxx*。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	Name *[]string `json:"name,omitempty"`

	// **参数解释**：规格类别。 支持多值查询，查询条件格式：*type=xxx&type=xxx*。  **约束限制**：不涉及  **取值范围**： - L4和L7 表示四层网络型和七层应用型flavor。 [- gateway 表示网关型LB的flavor，目前只支持弹性计费类型。当前仅支持欧洲局点。](tag:hws_eu) - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max、L7_elastic_max[和gateway_elastic_max](tag:hws_eu) 表示弹性扩缩容实例的上限规格。  **默认取值**：不涉及
	Type *[]string `json:"type,omitempty"`

	// **参数解释**：是否查询公共规格。  **约束限制**：不涉及  **取值范围**： - true表示查询公共规格，所有租户可见的规格。 - false表示查询私有规格，当前仅租户可见的规格。  **默认取值**：不涉及
	Shared *bool `json:"shared,omitempty"`

	// **参数解释**：公网边界组。 支持多值查询，查询条件格式：*public_border_group=xxx&public_border_group=xxx*。  **约束限制**：不涉及  **取值范围**： - center：表示中心站点的公网边界组 - 边缘站点名称：表示边缘站点的公网边界组  **默认取值**：不涉及
	PublicBorderGroup *[]string `json:"public_border_group,omitempty"`

	// **参数解释**：可用区子类型编码。该字段主要用于区分在边缘场景下，边缘AZ的类型。 支持多值查询，查询条件格式：*category=xxx&category=xxx*。  **约束限制**：不涉及  **取值范围**：0表示center，21表示homezone，41表示IES。  **默认取值**：不涉及
	Category *[]int32 `json:"category,omitempty"`

	// **参数解释**：是否查询当前租户所有的弹性上限规格。  **约束限制**：不涉及  **取值范围**： - true：查询当前租户所有的弹性上限规格（l4_elastic_max、l7_elastic_max）。 - false：查询租户弹性上限规格中最大的规格（l4类型优先比较cps指标，然后是带宽；l7类型优先比较https cps指标然后是qps指标）。  **默认取值**：不涉及
	ListAll *bool `json:"list_all,omitempty"`

	// **参数解释**：[是否售罄。](tag:hws)[是否无法购买该规格的LB。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,fcs,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg,ct)  **约束限制**：不涉及  **取值范围**： - true：[已售罄，将](tag:hws)无法购买该规格的LB。 - false：[未售罄，](tag:hws)可购买该规格的LB。  **默认取值**：不涉及
	FlavorSoldOut *bool `json:"flavor_sold_out,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
