package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeGroupsResponse Response Object
type DescribeGroupsResponse struct {

	// 用户组列表
	Groups         *[]DescribeGroupResp `json:"groups,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o DescribeGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeGroupsResponse struct{}"
	}

	return strings.Join([]string{"DescribeGroupsResponse", string(data)}, " ")
}
