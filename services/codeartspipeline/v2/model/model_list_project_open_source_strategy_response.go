package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectOpenSourceStrategyResponse Response Object
type ListProjectOpenSourceStrategyResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 开源治理策略列表
	Data           *[]SimpleOpenSourceRuleSetVo `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListProjectOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"ListProjectOpenSourceStrategyResponse", string(data)}, " ")
}
