package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersResponse Response Object
type ListImageMembersResponse struct {

	// 成员信息
	Members *[]ImageMember `json:"members,omitempty"`

	// 视图信息
	Schema *string `json:"schema,omitempty"`

	PageInfo       *GlancePageInfo `json:"page_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListImageMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersResponse struct{}"
	}

	return strings.Join([]string{"ListImageMembersResponse", string(data)}, " ")
}
