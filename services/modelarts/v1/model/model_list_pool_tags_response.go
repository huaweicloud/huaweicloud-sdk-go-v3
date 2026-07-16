package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolTagsResponse Response Object
type ListPoolTagsResponse struct {

	// **参数解释**：资源标签的列表。
	Tags           *[]PoolTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListPoolTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPoolTagsResponse", string(data)}, " ")
}
