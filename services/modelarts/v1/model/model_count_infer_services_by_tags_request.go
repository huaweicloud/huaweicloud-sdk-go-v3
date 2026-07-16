package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountInferServicesByTagsRequest Request Object
type CountInferServicesByTagsRequest struct {

	// **参数解释：** 工作空间ID，workspaceId将会被设置为null。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Body *QueryTmsResourceCountRequest `json:"body,omitempty"`
}

func (o CountInferServicesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountInferServicesByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountInferServicesByTagsRequest", string(data)}, " ")
}
