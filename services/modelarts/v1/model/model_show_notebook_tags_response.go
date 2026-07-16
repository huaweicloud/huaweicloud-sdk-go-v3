package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotebookTagsResponse Response Object
type ShowNotebookTagsResponse struct {

	// **参数解释**：标签的融合结构，相同key合并。
	Tags           *[]CombineTmsTags `json:"tags,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowNotebookTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotebookTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowNotebookTagsResponse", string(data)}, " ")
}
