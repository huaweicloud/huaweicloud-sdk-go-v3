package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InspectResult struct {

	// 数据类型。
	Type *string `json:"type,omitempty"`

	// 条目数。
	CountNum *int32 `json:"count_num,omitempty"`

	// 重复读。
	Multiplicity *float64 `json:"multiplicity,omitempty"`

	// 合法率。
	LegalRate *float64 `json:"legal_rate,omitempty"`
}

func (o InspectResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InspectResult struct{}"
	}

	return strings.Join([]string{"InspectResult", string(data)}, " ")
}
