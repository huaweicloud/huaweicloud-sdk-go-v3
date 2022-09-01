package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Flow struct {

	// 流程id。
	FlowId string `json:"flow_id" xml:"flow_id"`

	// 属性对过滤。
	AttrPairRulesFilter *[]AttrPairRules `json:"attr_pair_rules_filter,omitempty" xml:"attr_pair_rules_filter"`

	// 属性对保留。
	AttrPairRulesReserve *[]AttrPairRules `json:"attr_pair_rules_reserve,omitempty" xml:"attr_pair_rules_reserve"`

	// 属性去重。
	DeduplicationList *[]Deduplication `json:"deduplication_list,omitempty" xml:"deduplication_list"`

	AttributeInfo *AttributeInfo `json:"attribute_info,omitempty" xml:"attribute_info"`

	BloomFilterConf *BloomFilterConf `json:"bloom_filter_conf,omitempty" xml:"bloom_filter_conf"`

	// 分组打散属性。
	GroupAttr *string `json:"group_attr,omitempty" xml:"group_attr"`

	// 在排序前去重。
	PreDeal *bool `json:"pre_deal,omitempty" xml:"pre_deal"`

	// 排序配置信息。
	RankSetting *string `json:"rank_setting,omitempty" xml:"rank_setting"`

	Rules *Rule `json:"rules,omitempty" xml:"rules"`

	// 过滤配置信息。
	FilterSets *[]string `json:"filter_sets,omitempty" xml:"filter_sets"`

	// 属性值过滤。
	AttrValueRulesFilter *[]AttrValueRules `json:"attr_value_rules_filter,omitempty" xml:"attr_value_rules_filter"`

	// 属性值保留。
	AttrValueRulesReserve *[]AttrValueRules `json:"attr_value_rules_reserve,omitempty" xml:"attr_value_rules_reserve"`

	// 排序作业（使用点击率预估时需要提供此参数）。
	CtrJob *string `json:"ctr_job,omitempty" xml:"ctr_job"`

	// 流量占比。
	Ratio *int32 `json:"ratio,omitempty" xml:"ratio"`

	// 需要置顶的候选集列表。
	Toppings *[]string `json:"toppings,omitempty" xml:"toppings"`
}

func (o Flow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flow struct{}"
	}

	return strings.Join([]string{"Flow", string(data)}, " ")
}
