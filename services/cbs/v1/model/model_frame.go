package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Frame struct {

	// 意图
	Intention string `json:"intention" xml:"intention"`

	// 命中意图置信度。
	Confidence float64 `json:"confidence" xml:"confidence"`

	// 当前槽位列表。
	CurrentSlots []CurrentSlot `json:"current_slots" xml:"current_slots"`

	// 历史槽位列表。
	HistorySlots []HistorySlot `json:"history_slots" xml:"history_slots"`

	// 机器人回复。
	Reply string `json:"reply" xml:"reply"`

	// 任务是否完成。
	TaskComplete bool `json:"task_complete" xml:"task_complete"`

	// 对话流程是否完成。
	FlowComplete bool `json:"flow_complete" xml:"flow_complete"`

	// 候选词。
	CandidateWords *[]string `json:"candidate_words,omitempty" xml:"candidate_words"`

	// 意图名称
	IntentionAlias string `json:"intention_alias" xml:"intention_alias"`
}

func (o Frame) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Frame struct{}"
	}

	return strings.Join([]string{"Frame", string(data)}, " ")
}
