package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServicesByTagsRequest Request Object
type ListInferServicesByTagsRequest struct {

	// **参数解释：** 指定返回的最大条目数。 **约束限制：** 不涉及。 **取值范围：** [1,500] **默认取值：** 10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 分页列表查询的偏移量。 **约束限制：** offset必须是limit的整数倍。 **取值范围：** 不涉及。 **默认取值：** 0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 工作空间ID，workspaceId将会被设置为null。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Body *QueryTmsResourceRequest `json:"body,omitempty"`
}

func (o ListInferServicesByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServicesByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInferServicesByTagsRequest", string(data)}, " ")
}
