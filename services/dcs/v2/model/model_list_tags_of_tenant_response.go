package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTagsOfTenantResponse struct {

	// 标签列表。
	Tags           *[]Tag `json:"tags,omitempty" xml:"tags"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsOfTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsOfTenantResponse struct{}"
	}

	return strings.Join([]string{"ListTagsOfTenantResponse", string(data)}, " ")
}
