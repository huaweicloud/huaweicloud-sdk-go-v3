package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFlavorsRequest Request Object
type ListFlavorsRequest struct {

	// 参数解释：上一页最后一条记录的ID。  约束限制： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。
	Marker *string `json:"marker,omitempty"`

	// 参数解释：每页返回的个数。  取值范围：0-2000  默认取值：2000
	Limit *int32 `json:"limit,omitempty"`

	// 参数解释：是否反向查询。  约束限制： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  取值范围： - true：查询上一页。 - false：查询下一页，默认。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 参数解释：规格ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 参数解释：规格名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 参数解释：规格类别。  取值范围： - L4和L7 表示四层网络型和七层应用型flavor。 [- gateway 表示网关型LB的flavor，目前只支持弹性计费类型。当前仅支持欧洲局点。](tag:hws_eu) - L4_elastic和L7_elastic 表示弹性扩缩容实例的下限规格。 - L4_elastic_max、L7_elastic_max[和gateway_elastic_max](tag:hws_eu) 表示弹性扩缩容实例的上限规格。  支持多值查询，查询条件格式：*type=xxx&type=xxx*。
	Type *[]string `json:"type,omitempty"`

	// 参数解释：是否查询公共规格。  取值范围：true表示公共规格，所有租户可见。false表示私有规格，为当前租户所有。
	Shared *bool `json:"shared,omitempty"`

	// 站点pbg。  支持多值查询，查询条件格式：*public_border_group=xxx&public_border_group=xxx*。
	PublicBorderGroup *[]string `json:"public_border_group,omitempty"`

	// 站点category。  支持多值查询，查询条件格式：*category=xxx&category=xxx*。
	Category *[]int32 `json:"category,omitempty"`

	// 设置list_all=true时，返回当前租户所有的弹性上限规格（l4_elastic_max、l7__elastic_max）。否则只返回该租户弹性上限规格中最大的规格（l4类型优先比较cps指标，然后是带宽；l7类型优先比较https cps指标然后是qps指标）。
	ListAll *bool `json:"list_all,omitempty"`

	// 参数解释： [是否售罄。](tag:hws) [是否无法购买该规格的LB。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,fcs,dt,dt_test,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,srg,ct)  取值范围： - true：[已售罄，将](tag:hws)无法购买该规格的LB。 - false：[未售罄，](tag:hws)可购买该规格的LB。
	FlavorSoldOut *bool `json:"flavor_sold_out,omitempty"`
}

func (o ListFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorsRequest", string(data)}, " ")
}
