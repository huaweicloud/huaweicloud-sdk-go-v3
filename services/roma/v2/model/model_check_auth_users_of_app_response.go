package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckAuthUsersOfAppResponse struct {

	// 应用的总成员数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 应用ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 用户成员列表
	Users          *[]AppUsersUsers `json:"users,omitempty" xml:"users"`
	HttpStatusCode int              `json:"-"`
}

func (o CheckAuthUsersOfAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAuthUsersOfAppResponse struct{}"
	}

	return strings.Join([]string{"CheckAuthUsersOfAppResponse", string(data)}, " ")
}
