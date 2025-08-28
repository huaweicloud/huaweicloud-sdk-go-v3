package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstructionInfo 指令信息。
type InstructionInfo struct {

	// 指令ID。
	InstructionId *string `json:"instruction_id,omitempty"`

	// 指令名称。
	Name *string `json:"name,omitempty"`

	// 指令描述。
	Description *string `json:"description,omitempty"`

	// 操作指令。
	Instruction *string `json:"instruction,omitempty"`

	// 槽位列表
	SlotList *[]InstructionSlotInfo `json:"slot_list,omitempty"`

	// 指令回复话术列表
	ReplyWordsList *[]InstructionReplyWordsInfo `json:"reply_words_list,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o InstructionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstructionInfo struct{}"
	}

	return strings.Join([]string{"InstructionInfo", string(data)}, " ")
}
