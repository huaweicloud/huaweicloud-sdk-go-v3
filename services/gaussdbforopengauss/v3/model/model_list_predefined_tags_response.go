package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
