package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FeatureTransformation
type FeatureTransformation struct {
	Attr *Attribute `json:"attr,omitempty"`

	// 离散方法： - equal_distance_discrete，等距离散 - user_define_discrete，自定义离散 - normalize，归一化 - null，不离散
	DiscreteMethod *string `json:"discrete_method,omitempty"`

	// 具体处理参数。
	Params *interface{} `json:"params,omitempty"`
}

func (o FeatureTransformation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FeatureTransformation struct{}"
	}

	return strings.Join([]string{"FeatureTransformation", string(data)}, " ")
}
