package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteScriptRequest Request Object
type ExecuteScriptRequest struct {

	// 脚本UUID
	ScriptUuid string `json:"script_uuid"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`

	Body *ScriptExecuteModel `json:"body,omitempty"`
}

func (o ExecuteScriptRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScriptRequest struct{}"
	}

	return strings.Join([]string{"ExecuteScriptRequest", string(data)}, " ")
}
