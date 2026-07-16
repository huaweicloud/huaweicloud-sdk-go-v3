package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferServiceTagRequest Request Object
type DeleteInferServiceTagRequest struct {

	// **参数解释：** 待创建标签的资源ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ResourceId string `json:"resource_id"`

	// **参数解释：** 工作空间ID，workspaceId将会被设置为null。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Body *DeleteInferTmsTagsRequest `json:"body,omitempty"`
}

func (o DeleteInferServiceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferServiceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteInferServiceTagRequest", string(data)}, " ")
}
