package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleKnowledgeLibraryInfo 角色知识库基本信息。
type RoleKnowledgeLibraryInfo struct {
	KnowledgeType *KnowledgeTypeEnum `json:"knowledge_type"`

	// 知识库列表
	KnowledgeLibraryInfoList []KnowledgeLibraryBaseInfo `json:"knowledge_library_info_list"`
}

func (o RoleKnowledgeLibraryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleKnowledgeLibraryInfo struct{}"
	}

	return strings.Join([]string{"RoleKnowledgeLibraryInfo", string(data)}, " ")
}
