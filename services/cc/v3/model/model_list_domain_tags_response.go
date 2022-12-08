package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDomainTagsResponse struct {

	// 标签列表
	Tags *[]AggTag `json:"tags,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDomainTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDomainTagsResponse", string(data)}, " ")
}
