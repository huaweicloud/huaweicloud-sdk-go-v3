package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FinalReport 数据探索报告
type FinalReport struct {

	// 报告生成时间。
	GeneratedTime *string `json:"generated_time,omitempty"`

	// 宽表条目数，行为数据去重以后的数目。
	WideTableNum *int64 `json:"wide_table_num,omitempty"`

	// 用户齐全度，一条行为中的用户是否在产生这条行为的时候拥有画像。
	UserCompleteDegree *float64 `json:"user_complete_degree,omitempty"`

	// 物品齐全度，一条行为中的物品是否在这条行为产生的时候拥有画像。
	ItemCompleteDegree *float64 `json:"item_complete_degree,omitempty"`

	// 行为次数统计。
	BhvCount map[string]int64 `json:"bhv_count,omitempty"`

	// 用户数字类型特征统计。
	UserLongFeatureReport *[]NumFeatureReport `json:"user_long_feature_report,omitempty"`

	// 用户连续类型特征统计。
	UserFloatFeatureReport *[]NumFeatureReport `json:"user_float_feature_report,omitempty"`

	// 用户单值离散值类型特征统计。
	UserStrFeatureReport *[]StrFeatureReport `json:"user_str_feature_report,omitempty"`

	// 用户多值离散值类型特征统计。
	UserStrArrayFeatureReport *[]StrFeatureReport `json:"user_strArray_feature_report,omitempty"`

	// 物品数字类型特征统计。
	ItemLongFeatureReport *[]NumFeatureReport `json:"item_long_feature_report,omitempty"`

	// 物品连续类型特征统计。
	ItemFloatFeatureReport *[]NumFeatureReport `json:"item_float_feature_report,omitempty"`

	// 物品单值离散值类型特征统计。
	ItemStrFeatureReport *[]StrFeatureReport `json:"item_str_feature_report,omitempty"`

	// 物品多值离散值类型特征统计。
	ItemStrArrayFeatureReport *[]StrFeatureReport `json:"item_strArray_feature_report,omitempty"`
}

func (o FinalReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinalReport struct{}"
	}

	return strings.Join([]string{"FinalReport", string(data)}, " ")
}
