package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleBusinessReq 角色业务配置请求。
type RoleBusinessReq struct {
	Language *LanguageEnum `json:"language"`

	// 提示词。
	Prompt *string `json:"prompt,omitempty"`

	// 知识库列表
	KnowledgeLibraryList *[]RoleKnowledgeLibraryReq `json:"knowledge_library_list,omitempty"`
}

func (o RoleBusinessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleBusinessReq struct{}"
	}

	return strings.Join([]string{"RoleBusinessReq", string(data)}, " ")
}
