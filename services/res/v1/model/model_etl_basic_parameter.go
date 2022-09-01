package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type EtlBasicParameter struct {

	// 用户特征。
	UserFeatures *[]FeatureTransformation `json:"user_features,omitempty" xml:"user_features"`

	// 物品特征。
	ItemFeatures *[]FeatureTransformation `json:"item_features,omitempty" xml:"item_features"`

	RankEtlFilter *RankEtlFilter `json:"rank_etl_filter,omitempty" xml:"rank_etl_filter"`
}

func (o EtlBasicParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EtlBasicParameter struct{}"
	}

	return strings.Join([]string{"EtlBasicParameter", string(data)}, " ")
}
