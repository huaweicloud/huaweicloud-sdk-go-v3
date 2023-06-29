package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MatchFeaturePair
type MatchFeaturePair struct {

	// 用户特征。
	UserFeatureName *string `json:"user_feature_name,omitempty"`

	// 物品特征。
	ItemFeatureName *string `json:"item_feature_name,omitempty"`

	// 权重。
	Weight *float64 `json:"weight,omitempty"`

	// 匹配个数度量。
	MatchCount *bool `json:"match_count,omitempty"`
}

func (o MatchFeaturePair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchFeaturePair struct{}"
	}

	return strings.Join([]string{"MatchFeaturePair", string(data)}, " ")
}
