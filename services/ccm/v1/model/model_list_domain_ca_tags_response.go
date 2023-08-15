package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainCaTagsResponse Response Object
type ListDomainCaTagsResponse struct {

	// 标签列表。
	Tags           *[]DomainTags `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDomainCaTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainCaTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainCaTagsResponse", string(data)}, " ")
}
