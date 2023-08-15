package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPredefinedTagsResponse Response Object
type ListPredefinedTagsResponse struct {
	Tags           *[][]interface{} `json:"tags,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListPredefinedTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPredefinedTagsResponse struct{}"
	}

	return strings.Join([]string{"ListPredefinedTagsResponse", string(data)}, " ")
}
