package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttackSubItem 攻击子类型分布项
type AttackSubItem struct {

	// 子类型标识： 合规检测子类型：politics_matched, violent_matched, illegal_matched, hatespeech_matched, privatesensitive_matched, religion_matched, other_matched 注入检测子类型：goal_hijacking_matched, prompt_leaking_matched, role_play_instruction_matched, unsafe_instruction_topic_matched, inquiry_with_unsafe_opinion_matched, reverse_exposure_matched, multi_step_instruction_matched, context_task_matched, opposition_response_matched, virtual_conversation_matched, do_anything_now_matched, code_injection_matched
	Key string `json:"key"`

	// 该子类型命中次数
	Num int64 `json:"num"`
}

func (o AttackSubItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttackSubItem struct{}"
	}

	return strings.Join([]string{"AttackSubItem", string(data)}, " ")
}
