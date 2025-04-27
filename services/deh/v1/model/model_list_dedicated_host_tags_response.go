package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDedicatedHostTagsResponse Response Object
type ListDedicatedHostTagsResponse struct {

	// 全部专属主机标签列表。
	Tags           *[]ListResourceTag `json:"tags,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListDedicatedHostTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedHostTagsResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedHostTagsResponse", string(data)}, " ")
}
