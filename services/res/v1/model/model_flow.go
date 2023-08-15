package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Flow
type Flow struct {

	// 流程id。
	FlowId string `json:"flow_id"`

	// 属性对过滤。
	AttrPairRulesFilter *[]AttrPairRules `json:"attr_pair_rules_filter,omitempty"`

	// 属性对保留。
	AttrPairRulesReserve *[]AttrPairRules `json:"attr_pair_rules_reserve,omitempty"`

	// 属性去重。
	DeduplicationList *[]Deduplication `json:"deduplication_list,omitempty"`

	AttributeInfo *AttributeInfo `json:"attribute_info,omitempty"`

	BloomFilterConf *BloomFilterConf `json:"bloom_filter_conf,omitempty"`

	// 分组打散属性。
	GroupAttr *string `json:"group_attr,omitempty"`

	// 在排序前去重。
	PreDeal *bool `json:"pre_deal,omitempty"`

	// 排序配置信息。
	RankSetting *string `json:"rank_setting,omitempty"`

	Rules *Rule `json:"rules,omitempty"`

	// 过滤配置信息。
	FilterSets *[]string `json:"filter_sets,omitempty"`

	// 属性值过滤。
	AttrValueRulesFilter *[]AttrValueRules `json:"attr_value_rules_filter,omitempty"`

	// 属性值保留。
	AttrValueRulesReserve *[]AttrValueRules `json:"attr_value_rules_reserve,omitempty"`

	// 排序作业（使用点击率预估时需要提供此参数）。
	CtrJob *string `json:"ctr_job,omitempty"`

	// 流量占比。
	Ratio *int32 `json:"ratio,omitempty"`

	// 需要置顶的候选集列表。
	Toppings *[]string `json:"toppings,omitempty"`
}

func (o Flow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Flow struct{}"
	}

	return strings.Join([]string{"Flow", string(data)}, " ")
}
