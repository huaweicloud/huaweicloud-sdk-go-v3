package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SyncIamUsersResponse struct {

	// 创建成功用户列表
	SyncUser       *[]string `json:"sync_user,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SyncIamUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncIamUsersResponse struct{}"
	}

	return strings.Join([]string{"SyncIamUsersResponse", string(data)}, " ")
}
