package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProtectedInstanceTagsResponse struct {
	// 标签列表。

	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListProtectedInstanceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstanceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedInstanceTagsResponse", string(data)}, " ")
}
