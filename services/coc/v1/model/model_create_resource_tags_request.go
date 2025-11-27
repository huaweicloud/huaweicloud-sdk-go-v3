package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagsRequest Request Object
type CreateResourceTagsRequest struct {

	// **参数解释：** 资源ID。 **约束限制：** 不涉及。 **取值范围：** 标签对应的资源id。 **默认取值：** 不涉及。
	ResourceId string `json:"resource_id"`

	Body *ResourceTagOperateRequest `json:"body,omitempty"`
}

func (o CreateResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceTagsRequest", string(data)}, " ")
}
