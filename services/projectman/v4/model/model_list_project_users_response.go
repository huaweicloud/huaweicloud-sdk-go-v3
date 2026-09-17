package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectUsersResponse Response Object
type ListProjectUsersResponse struct {

	// 返回信息
	Message *string `json:"message,omitempty"`

	// 返回用户列表
	Result *[]UserVo `json:"result,omitempty"`

	// 返回状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectUsersResponse struct{}"
	}

	return strings.Join([]string{"ListProjectUsersResponse", string(data)}, " ")
}
