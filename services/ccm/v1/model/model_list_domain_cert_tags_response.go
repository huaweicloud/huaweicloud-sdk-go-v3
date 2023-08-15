package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainCertTagsResponse Response Object
type ListDomainCertTagsResponse struct {

	// 标签列表。
	Tags           *[]DomainTags `json:"tags,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDomainCertTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainCertTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainCertTagsResponse", string(data)}, " ")
}
