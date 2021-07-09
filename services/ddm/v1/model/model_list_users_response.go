package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListUsersResponse struct {
	// DDM实例帐号相关信息的集合。

	Users *[]GetUsersListDetailResponses `json:"users,omitempty"`
	// 当前页码

	PageNo *int32 `json:"page_no,omitempty"`
	// 当前页码的数据条数

	PageSize *int32 `json:"page_size,omitempty"`
	// 总条数

	TotalRecord *int32 `json:"total_record,omitempty"`
	// 总页数

	TotalPage      *int32 `json:"total_page,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
