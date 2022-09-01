package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUsersResponse struct {

	// DDM实例帐号相关信息的集合。
	Users *[]GetUsersListDetailResponses `json:"users,omitempty" xml:"users"`

	// 当前页码
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no"`

	// 当前页码的数据条数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 总条数
	TotalRecord *int32 `json:"total_record,omitempty" xml:"total_record"`

	// 总页数
	TotalPage      *int32 `json:"total_page,omitempty" xml:"total_page"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
