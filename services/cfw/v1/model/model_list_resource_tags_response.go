package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTagsResponse Response Object
type ListResourceTagsResponse struct {
	Data *string `json:"data,omitempty"`

	// 资源标签列表
	Tags *[]ResourceTag `json:"tags,omitempty"`

	SysTags        *[]ResourceTag `json:"sys_tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ListResourceTagsResponse", string(data)}, " ")
}
