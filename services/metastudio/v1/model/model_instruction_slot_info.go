package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstructionSlotInfo 指令槽位信息。
type InstructionSlotInfo struct {

	// 槽位名称。
	Name string `json:"name"`

	// 槽位描述。
	Description *string `json:"description,omitempty"`

	// 知识库列表
	KnowledgeLibraryList *[]SlotKnowledgeLibraryInfo `json:"knowledge_library_list,omitempty"`
}

func (o InstructionSlotInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstructionSlotInfo struct{}"
	}

	return strings.Join([]string{"InstructionSlotInfo", string(data)}, " ")
}
