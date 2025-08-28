package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleBusinessInfo 角色业务配置基本信息。
type RoleBusinessInfo struct {

	// 角色业务配置ID。
	RoleBusinessId *string `json:"role_business_id,omitempty"`

	// 角色ID。
	RoleId *string `json:"role_id,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 提示词。
	Prompt *string `json:"prompt,omitempty"`

	// 知识库列表
	KnowledgeLibraryList *[]RoleKnowledgeLibraryInfo `json:"knowledge_library_list,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o RoleBusinessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleBusinessInfo struct{}"
	}

	return strings.Join([]string{"RoleBusinessInfo", string(data)}, " ")
}
