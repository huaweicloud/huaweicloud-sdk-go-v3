package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptResourceTagsResponse Response Object
type ListScriptResourceTagsResponse struct {
	Data           *ListTagsResponse `json:"data,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListScriptResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptResourceTagsResponse", string(data)}, " ")
}
