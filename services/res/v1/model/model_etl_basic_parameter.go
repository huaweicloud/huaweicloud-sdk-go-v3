package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EtlBasicParameter
type EtlBasicParameter struct {

	// 用户特征。
	UserFeatures *[]FeatureTransformation `json:"user_features,omitempty"`

	// 物品特征。
	ItemFeatures *[]FeatureTransformation `json:"item_features,omitempty"`

	RankEtlFilter *RankEtlFilter `json:"rank_etl_filter,omitempty"`
}

func (o EtlBasicParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EtlBasicParameter struct{}"
	}

	return strings.Join([]string{"EtlBasicParameter", string(data)}, " ")
}
