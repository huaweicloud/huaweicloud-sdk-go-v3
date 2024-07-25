package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsResponse Response Object
type ShowResourceTagsResponse struct {

	// 资源标签
	Tags *[]ResourceTagBody `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]ResourceTagBody `json:"sys_tags,omitempty"`

	// 执行动作
	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResourceTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsResponse", string(data)}, " ")
}
