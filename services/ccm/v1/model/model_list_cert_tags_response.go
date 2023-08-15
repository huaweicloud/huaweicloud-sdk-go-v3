package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertTagsResponse Response Object
type ListCertTagsResponse struct {

	// 标签列表，key和value键值对的集合。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCertTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertTagsResponse struct{}"
	}

	return strings.Join([]string{"ListCertTagsResponse", string(data)}, " ")
}
