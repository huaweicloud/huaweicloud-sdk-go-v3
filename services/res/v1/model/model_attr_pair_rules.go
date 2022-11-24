package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 属性对规则
type AttrPairRules struct {

	// 属性对。
	AttrPairs *[]AttrPair `json:"attr_pairs,omitempty"`
}

func (o AttrPairRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttrPairRules struct{}"
	}

	return strings.Join([]string{"AttrPairRules", string(data)}, " ")
}
