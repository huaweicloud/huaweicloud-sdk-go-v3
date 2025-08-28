package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstructionResponse Response Object
type ShowInstructionResponse struct {

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

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstructionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstructionResponse struct{}"
	}

	return strings.Join([]string{"ShowInstructionResponse", string(data)}, " ")
}
