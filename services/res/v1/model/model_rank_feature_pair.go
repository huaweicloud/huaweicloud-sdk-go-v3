package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type RankFeaturePair struct {

	// 待推荐对象的属性。
	FeatureNameA *string `json:"feature_name_a,omitempty" xml:"feature_name_a"`

	// 被推荐对象的属性。
	FeatureNameB *string `json:"feature_name_b,omitempty" xml:"feature_name_b"`

	// 权重。
	Weight *float32 `json:"weight,omitempty" xml:"weight"`
}

func (o RankFeaturePair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RankFeaturePair struct{}"
	}

	return strings.Join([]string{"RankFeaturePair", string(data)}, " ")
}
