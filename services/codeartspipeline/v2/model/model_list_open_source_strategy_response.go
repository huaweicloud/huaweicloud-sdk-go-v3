package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOpenSourceStrategyResponse Response Object
type ListOpenSourceStrategyResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 开源治理策略列表
	Data           *[]SimpleOpenSourceRuleSetVo `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"ListOpenSourceStrategyResponse", string(data)}, " ")
}
