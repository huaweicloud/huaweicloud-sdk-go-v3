package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataspaceRequest Request Object
type CreateDataspaceRequest struct {

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	Body *CreateDataspaceRequestBody `json:"body,omitempty"`
}

func (o CreateDataspaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataspaceRequest struct{}"
	}

	return strings.Join([]string{"CreateDataspaceRequest", string(data)}, " ")
}
