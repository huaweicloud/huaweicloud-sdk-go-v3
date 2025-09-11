package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditTagsResponse Response Object
type ListAuditTagsResponse struct {

	// 自定义标签
	Tags *[]ResourceTagInfoBean `json:"tags,omitempty"`

	// 系统标签
	SysTags        *[]ResourceTagInfoBean `json:"sys_tags,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAuditTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditTagsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditTagsResponse", string(data)}, " ")
}
