package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SyncIamUsersResponse Response Object
type SyncIamUsersResponse struct {

	// **参数解释**： 创建成功用户列表。 **取值范围**： 不涉及。
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
