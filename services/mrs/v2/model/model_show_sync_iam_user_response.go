package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSyncIamUserResponse Response Object
type ShowSyncIamUserResponse struct {

	// 已经同步了的用户
	UserNames *[]string `json:"user_names,omitempty"`

	// 已经同步了的用户组
	GroupNames     *[]string `json:"group_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowSyncIamUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSyncIamUserResponse struct{}"
	}

	return strings.Join([]string{"ShowSyncIamUserResponse", string(data)}, " ")
}
