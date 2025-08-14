package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeGroupsRequestBody 查询用户组信息请求体
type DescribeGroupsRequestBody struct {

	// 用户组唯一标识符（ID）列表
	GroupIds []string `json:"group_ids"`
}

func (o DescribeGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"DescribeGroupsRequestBody", string(data)}, " ")
}
