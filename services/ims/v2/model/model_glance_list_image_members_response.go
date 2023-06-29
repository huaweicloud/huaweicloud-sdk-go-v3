package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceListImageMembersResponse Response Object
type GlanceListImageMembersResponse struct {

	// 成员信息
	Members *[]GlanceImageMembers `json:"members,omitempty"`

	// 视图信息
	Schema         *string `json:"schema,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o GlanceListImageMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceListImageMembersResponse struct{}"
	}

	return strings.Join([]string{"GlanceListImageMembersResponse", string(data)}, " ")
}
