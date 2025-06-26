package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubResourceTagsResponse Response Object
type ListSubResourceTagsResponse struct {

	// 资源标签列表
	Tags           *[]ResourceTag `json:"tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListSubResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListSubResourceTagsResponse", string(data)}, " ")
}
