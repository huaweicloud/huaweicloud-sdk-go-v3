package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagRequest **参数解释**：删除资源标签结构体，支持批量删除。
type DeleteTagRequest struct {

	// **参数解释**：要删除的标签列表。 **约束限制**：不涉及。
	Tags []DeleteTagItem `json:"tags"`
}

func (o DeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
