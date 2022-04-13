package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddUserToAppResponse struct {
	// 应用的总成员数量

	Total *int32 `json:"total,omitempty"`
	// 应用ID

	Id *string `json:"id,omitempty"`
	// 用户成员列表

	Users          *[]AppUsersUsers `json:"users,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o AddUserToAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddUserToAppResponse struct{}"
	}

	return strings.Join([]string{"AddUserToAppResponse", string(data)}, " ")
}
