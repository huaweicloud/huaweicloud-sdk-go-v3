package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClusterTagsResponse struct {

	// 标签列表。
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListClusterTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClusterTagsResponse", string(data)}, " ")
}
