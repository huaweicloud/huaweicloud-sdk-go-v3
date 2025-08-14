package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeUsersRequestBody 查询用户信息请求体
type DescribeUsersRequestBody struct {

	// 用户唯一标识符（ID）列表
	UserIds []string `json:"user_ids"`
}

func (o DescribeUsersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeUsersRequestBody struct{}"
	}

	return strings.Join([]string{"DescribeUsersRequestBody", string(data)}, " ")
}
