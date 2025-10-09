package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetResourceTagsResponse Response Object
type GetResourceTagsResponse struct {

	// **参数解释**： 自定义标签 **约束限制**： 不涉及
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// **参数解释**： 系统标签 **约束限制**： 不涉及
	SysTags        *[]ResourceTag `json:"sys_tags,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o GetResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"GetResourceTagsResponse", string(data)}, " ")
}
