package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRecordRuleRequest Request Object
type CreateRecordRuleRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 应用id
	AppId string `json:"app_id"`

	Body *RecordRuleReq `json:"body,omitempty"`
}

func (o CreateRecordRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordRuleRequest", string(data)}, " ")
}
