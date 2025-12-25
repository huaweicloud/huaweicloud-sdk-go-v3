package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMapperRequest Request Object
type CreateMapperRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *CreateDpeMappingRequestBody `json:"body,omitempty"`
}

func (o CreateMapperRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMapperRequest struct{}"
	}

	return strings.Join([]string{"CreateMapperRequest", string(data)}, " ")
}
