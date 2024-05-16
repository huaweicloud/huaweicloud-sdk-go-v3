package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportKnowledgeSkillRequest Request Object
type ExportKnowledgeSkillRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 技能ID。
	SkillId string `json:"skill_id"`

	// 导出格式类型。0：科大讯飞
	ExportType int32 `json:"export_type"`
}

func (o ExportKnowledgeSkillRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportKnowledgeSkillRequest struct{}"
	}

	return strings.Join([]string{"ExportKnowledgeSkillRequest", string(data)}, " ")
}
