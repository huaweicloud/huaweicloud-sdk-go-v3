package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConformancePackScore 合规规则包的分数详情。
type ConformancePackScore struct {

	// 合规规则包ID。
	Id *string `json:"id,omitempty"`

	// 合规规则包名称。
	Name *string `json:"name,omitempty"`

	// 合规规则包分数。
	Score *string `json:"score,omitempty"`
}

func (o ConformancePackScore) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConformancePackScore struct{}"
	}

	return strings.Join([]string{"ConformancePackScore", string(data)}, " ")
}
