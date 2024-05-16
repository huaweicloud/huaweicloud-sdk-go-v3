package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeIntentRequest Request Object
type CreateKnowledgeIntentRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *CreateKnowledgeIntentReq `json:"body,omitempty"`
}

func (o CreateKnowledgeIntentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeIntentRequest struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeIntentRequest", string(data)}, " ")
}
