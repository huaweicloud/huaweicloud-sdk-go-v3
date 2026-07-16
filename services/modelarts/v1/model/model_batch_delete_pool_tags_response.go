package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoolTagsResponse Response Object
type BatchDeletePoolTagsResponse struct {

	// **参数解释**：资源标签的列表。
	Tags           *[]PoolTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o BatchDeletePoolTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoolTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePoolTagsResponse", string(data)}, " ")
}
