package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 物品离散值类型特征统计
type StrFeatureReport struct {

	// 特征名。
	Name *string `json:"name,omitempty" xml:"name"`

	// 特征类型。
	DataType *string `json:"data_type,omitempty" xml:"data_type"`

	// 离散类型特征出现次数统计。
	StrCount map[string]int32 `json:"str_count,omitempty" xml:"str_count"`
}

func (o StrFeatureReport) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StrFeatureReport struct{}"
	}

	return strings.Join([]string{"StrFeatureReport", string(data)}, " ")
}
