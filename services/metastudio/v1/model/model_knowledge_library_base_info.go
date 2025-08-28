package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeLibraryBaseInfo 知识库基础信息。
type KnowledgeLibraryBaseInfo struct {

	// 知识库ID。
	KnowledgeLibraryId string `json:"knowledge_library_id"`

	// 知识库名称。
	Name *string `json:"name,omitempty"`
}

func (o KnowledgeLibraryBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeLibraryBaseInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeLibraryBaseInfo", string(data)}, " ")
}
