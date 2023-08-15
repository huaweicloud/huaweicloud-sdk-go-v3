package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RankFeaturePair
type RankFeaturePair struct {

	// 待推荐对象的属性。
	FeatureNameA *string `json:"feature_name_a,omitempty"`

	// 被推荐对象的属性。
	FeatureNameB *string `json:"feature_name_b,omitempty"`

	// 权重。
	Weight *float32 `json:"weight,omitempty"`
}

func (o RankFeaturePair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RankFeaturePair struct{}"
	}

	return strings.Join([]string{"RankFeaturePair", string(data)}, " ")
}
