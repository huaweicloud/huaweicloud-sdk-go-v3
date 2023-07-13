package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectTagsListResponse Response Object
type ShowProjectTagsListResponse struct {

	// 标签列表
	Tags *[]TagItem `json:"tags,omitempty"`

	// 系统标签列表
	SysTags        *[]TagItem `json:"sys_tags,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowProjectTagsListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectTagsListResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectTagsListResponse", string(data)}, " ")
}
