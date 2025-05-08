package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptScriptRequest Request Object
type AcceptScriptRequest struct {

	// 脚本Uuid
	ScriptUuid string `json:"script_uuid"`

	// 项目ID。
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息。
	XUserProfile *string `json:"x-user-profile,omitempty"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ApprovalJobScriptModel `json:"body,omitempty"`
}

func (o AcceptScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScriptRequest struct{}"
	}

	return strings.Join([]string{"AcceptScriptRequest", string(data)}, " ")
}
