package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDbUsersResponse Response Object
type ListDbUsersResponse struct {

	// 列表中每个元素表示一个数据库用户。
	Users *[]GaussDBforOpenGaussUserForList `json:"users,omitempty"`

	// 数据库用户总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListDbUsersResponse", string(data)}, " ")
}
