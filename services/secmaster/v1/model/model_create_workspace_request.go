package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkspaceRequest Request Object
type CreateWorkspaceRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	Body *CreateWorkspaceRequestBody `json:"body,omitempty"`
}

func (o CreateWorkspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceRequest", string(data)}, " ")
}
