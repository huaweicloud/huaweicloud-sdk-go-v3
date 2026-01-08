package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListenerTagsResponse Response Object
type ShowListenerTagsResponse struct {

	// **参数解释**：标签列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowListenerTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerTagsResponse", string(data)}, " ")
}
