package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWatermarkRuleResponse Response Object
type ListWatermarkRuleResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 响应数据
	RulesInfos *[]WatermarkRule `json:"rules_infos,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListWatermarkRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWatermarkRuleResponse struct{}"
	}

	return strings.Join([]string{"ListWatermarkRuleResponse", string(data)}, " ")
}
