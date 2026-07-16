package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferTmsTagsRequest 删除资源标签结构体，支持批量删除。
type DeleteInferTmsTagsRequest struct {

	// **参数解释：** 要删除的标签列表。
	Tags []TmsTagForDeletion `json:"tags"`

	// **参数解释：** 待删除标签的资源ID。 **取值范围：** 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o DeleteInferTmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferTmsTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteInferTmsTagsRequest", string(data)}, " ")
}
