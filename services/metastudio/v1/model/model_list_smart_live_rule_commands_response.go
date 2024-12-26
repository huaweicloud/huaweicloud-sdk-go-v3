package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmartLiveRuleCommandsResponse Response Object
type ListSmartLiveRuleCommandsResponse struct {

	// 数字人直播任务互动规则未确认命令总数。
	Count *int32 `json:"count,omitempty"`

	// 数字人互动规则命令列表。
	RuleCommands *[]RuleCommand `json:"rule_commands,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSmartLiveRuleCommandsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmartLiveRuleCommandsResponse struct{}"
	}

	return strings.Join([]string{"ListSmartLiveRuleCommandsResponse", string(data)}, " ")
}
