package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreatePoolTagsResponse Response Object
type BatchCreatePoolTagsResponse struct {

	// **参数解释**：资源标签的列表。
	Tags           *[]PoolTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o BatchCreatePoolTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreatePoolTagsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreatePoolTagsResponse", string(data)}, " ")
}
