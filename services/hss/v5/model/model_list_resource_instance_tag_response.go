package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceInstanceTagResponse Response Object
type ListResourceInstanceTagResponse struct {

	// 标签对象列表
	Tags *[]ResourceTagInfo `json:"tags,omitempty"`

	// 系统标签对象列表
	SysTags        *[]ResourceTagInfo `json:"sys_tags,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListResourceInstanceTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceInstanceTagResponse struct{}"
	}

	return strings.Join([]string{"ListResourceInstanceTagResponse", string(data)}, " ")
}
