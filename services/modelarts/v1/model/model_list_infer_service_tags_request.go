package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServiceTagsRequest Request Object
type ListInferServiceTagsRequest struct {

	// **参数解释：** 工作空间ID，workspaceId将会被设置为null。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ListInferServiceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServiceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInferServiceTagsRequest", string(data)}, " ")
}
