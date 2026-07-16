package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPoolTagsResponse Response Object
type ShowPoolTagsResponse struct {

	// **参数解释**：资源标签的列表。
	Tags           *[]PoolTag `json:"tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowPoolTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPoolTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowPoolTagsResponse", string(data)}, " ")
}
