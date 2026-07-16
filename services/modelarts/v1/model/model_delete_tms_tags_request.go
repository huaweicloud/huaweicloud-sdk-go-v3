package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTmsTagsRequest 删除资源标签结构体，支持批量删除。
type DeleteTmsTagsRequest struct {

	// **参数解释**：要删除的标签列表。 **约束限制**：不涉及。
	Tags []TmsTag `json:"tags"`
}

func (o DeleteTmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTmsTagsRequest struct{}"
	}

	return strings.Join([]string{"DeleteTmsTagsRequest", string(data)}, " ")
}
