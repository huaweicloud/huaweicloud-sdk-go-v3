package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstructionReq 修改指令请求。
type UpdateInstructionReq struct {

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
}

func (o UpdateInstructionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstructionReq struct{}"
	}

	return strings.Join([]string{"UpdateInstructionReq", string(data)}, " ")
}
