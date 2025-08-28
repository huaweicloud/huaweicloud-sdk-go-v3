package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstructionReq 创建指令请求。
type CreateInstructionReq struct {

	// 指令名称。
	Name string `json:"name"`

	// 指令描述。
	Description string `json:"description"`

	// 操作指令。
	Instruction string `json:"instruction"`

	// 指令集ID。
	InstructionLibraryId string `json:"instruction_library_id"`

	// 槽位列表
	SlotList *[]InstructionSlotInfo `json:"slot_list,omitempty"`

	// 指令回复话术列表
	ReplyWordsList *[]InstructionReplyWordsInfo `json:"reply_words_list,omitempty"`
}

func (o CreateInstructionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstructionReq struct{}"
	}

	return strings.Join([]string{"CreateInstructionReq", string(data)}, " ")
}
