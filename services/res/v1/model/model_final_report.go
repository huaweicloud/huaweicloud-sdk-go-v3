package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据探索报告
type FinalReport struct {

	// 报告生成时间。
	GeneratedTime *string `json:"generated_time,omitempty" xml:"generated_time"`

	// 宽表条目数，行为数据去重以后的数目。
	WideTableNum *int64 `json:"wide_table_num,omitempty" xml:"wide_table_num"`

	// 用户齐全度，一条行为中的用户是否在产生这条行为的时候拥有画像。
	UserCompleteDegree *float64 `json:"user_complete_degree,omitempty" xml:"user_complete_degree"`

	// 物品齐全度，一条行为中的物品是否在这条行为产生的时候拥有画像。
	ItemCompleteDegree *float64 `json:"item_complete_degree,omitempty" xml:"item_complete_degree"`

	// 行为次数统计。
	BhvCount map[string]int64 `json:"bhv_count,omitempty" xml:"bhv_count"`

	// 用户数字类型特征统计。
	UserLongFeatureReport *[]NumFeatureReport `json:"user_long_feature_report,omitempty" xml:"user_long_feature_report"`

	// 用户连续类型特征统计。
	UserFloatFeatureReport *[]NumFeatureReport `json:"user_float_feature_report,omitempty" xml:"user_float_feature_report"`

	// 用户单值离散值类型特征统计。
	UserStrFeatureReport *[]StrFeatureReport `json:"user_str_feature_report,omitempty" xml:"user_str_feature_report"`

	// 用户多值离散值类型特征统计。
	UserStrArrayFeatureReport *[]StrFeatureReport `json:"user_strArray_feature_report,omitempty" xml:"user_strArray_feature_report"`

	// 物品数字类型特征统计。
	ItemLongFeatureReport *[]NumFeatureReport `json:"item_long_feature_report,omitempty" xml:"item_long_feature_report"`

	// 物品连续类型特征统计。
	ItemFloatFeatureReport *[]NumFeatureReport `json:"item_float_feature_report,omitempty" xml:"item_float_feature_report"`

	// 物品单值离散值类型特征统计。
	ItemStrFeatureReport *[]StrFeatureReport `json:"item_str_feature_report,omitempty" xml:"item_str_feature_report"`

	// 物品多值离散值类型特征统计。
	ItemStrArrayFeatureReport *[]StrFeatureReport `json:"item_strArray_feature_report,omitempty" xml:"item_strArray_feature_report"`
}

func (o FinalReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinalReport struct{}"
	}

	return strings.Join([]string{"FinalReport", string(data)}, " ")
}
