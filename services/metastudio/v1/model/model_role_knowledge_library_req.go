package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleKnowledgeLibraryReq 角色知识库配置请求信息。
type RoleKnowledgeLibraryReq struct {
	KnowledgeType *KnowledgeTypeEnum `json:"knowledge_type"`

	// 知识库ID列表
	KnowledgeLibraryIdList []string `json:"knowledge_library_id_list"`
}

func (o RoleKnowledgeLibraryReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleKnowledgeLibraryReq struct{}"
	}

	return strings.Join([]string{"RoleKnowledgeLibraryReq", string(data)}, " ")
}
