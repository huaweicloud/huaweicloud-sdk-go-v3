package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DivisionPattern struct {

	// 阶梯类型，1：一次分档模式、2：累计分档模式
	PatternType *int32 `json:"pattern_type,omitempty"`

	// 分档类型，1：STEP，按档拆分定价、2：TIER，按最终档定价
	DivisionType *int32 `json:"division_type,omitempty"`

	// 阶梯分档列表
	DivisionList *[]Division `json:"division_list,omitempty"`
}

func (o DivisionPattern) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DivisionPattern struct{}"
	}

	return strings.Join([]string{"DivisionPattern", string(data)}, " ")
}
