package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DtAggregateRequest 聚合计算
type DtAggregateRequest struct {

	// 输入参数，最多支持10个
	Inputs *[]InputRequest `json:"inputs,omitempty"`

	Outputs *[]OutputRequest `json:"outputs,omitempty"`
}

func (o DtAggregateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DtAggregateRequest struct{}"
	}

	return strings.Join([]string{"DtAggregateRequest", string(data)}, " ")
}
