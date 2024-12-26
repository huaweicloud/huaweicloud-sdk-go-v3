package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleCommand 数字人互动规则命令信息。
type RuleCommand struct {

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 命令ID。
	CommandId *string `json:"command_id,omitempty"`

	// 互动规则列表
	InteractionRules *[]LiveRoomInteractionRuleInfo `json:"interaction_rules,omitempty"`
}

func (o RuleCommand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleCommand struct{}"
	}

	return strings.Join([]string{"RuleCommand", string(data)}, " ")
}
