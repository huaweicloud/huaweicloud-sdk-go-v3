package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateScriptRequest Request Object
type CreateScriptRequest struct {

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	Body *AddScriptModel `json:"body,omitempty"`
}

func (o CreateScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScriptRequest struct{}"
	}

	return strings.Join([]string{"CreateScriptRequest", string(data)}, " ")
}
