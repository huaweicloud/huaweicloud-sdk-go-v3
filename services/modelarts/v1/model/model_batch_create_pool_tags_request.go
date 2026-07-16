package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePoolTagsRequest Request Object
type BatchCreatePoolTagsRequest struct {

	// **参数解释**：资源池名称。取自资源池详情的metadata字段中的name的值。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`

	Body *CreateTagRequest `json:"body,omitempty"`
}

func (o BatchCreatePoolTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePoolTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreatePoolTagsRequest", string(data)}, " ")
}
