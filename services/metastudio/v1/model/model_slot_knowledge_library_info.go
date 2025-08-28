package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlotKnowledgeLibraryInfo 槽位关联的知识库信息。
type SlotKnowledgeLibraryInfo struct {
	Language *LanguageEnum `json:"language"`

	// 知识库ID。
	KnowledgeLibraryId string `json:"knowledge_library_id"`
}

func (o SlotKnowledgeLibraryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlotKnowledgeLibraryInfo struct{}"
	}

	return strings.Join([]string{"SlotKnowledgeLibraryInfo", string(data)}, " ")
}
