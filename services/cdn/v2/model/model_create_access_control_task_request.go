package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccessControlTaskRequest Request Object
type CreateAccessControlTaskRequest struct {

	// **参数解释：** 操作类型， **约束限制：** 不涉及 **取值范围：** - ban：禁用 - unban：解禁  **默认取值：** 不涉及
	Action string `json:"action"`

	Body *UrlAccessControlTaskRequestBody `json:"body,omitempty"`
}

func (o CreateAccessControlTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccessControlTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAccessControlTaskRequest", string(data)}, " ")
}
