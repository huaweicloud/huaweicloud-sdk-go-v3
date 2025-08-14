package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeUsersResponse Response Object
type DescribeUsersResponse struct {

	// 用户详情列表
	Users          *[]DescribeUserResp `json:"users,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o DescribeUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeUsersResponse struct{}"
	}

	return strings.Join([]string{"DescribeUsersResponse", string(data)}, " ")
}
