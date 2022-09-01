package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProjectMembersV4Response struct {

	// 项目成员列表
	Members *[]MemberListV4Members `json:"members,omitempty" xml:"members"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProjectMembersV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectMembersV4Response struct{}"
	}

	return strings.Join([]string{"ListProjectMembersV4Response", string(data)}, " ")
}
