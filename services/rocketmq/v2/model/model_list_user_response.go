package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListUserResponse struct {

	// 用户列表。
	Users *[]User `json:"users,omitempty" xml:"users"`

	// 总用户个数。
	Total          float32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int     `json:"-"`
}

func (o ListUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserResponse struct{}"
	}

	return strings.Join([]string{"ListUserResponse", string(data)}, " ")
}
