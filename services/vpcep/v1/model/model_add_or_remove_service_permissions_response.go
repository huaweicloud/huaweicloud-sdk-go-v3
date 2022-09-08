package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddOrRemoveServicePermissionsResponse struct {

	// permission列表。 权限格式为iam:domain:: 6e9dfd51d1124e8d8498dce894923a0d或“*”， “*”表示所有用户的终端节点可连接。其中 6e9dfd51d1124e8d8498dce894923a0d为可连接 的用户domian_id。
	Permissions    *[]string `json:"permissions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddOrRemoveServicePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrRemoveServicePermissionsResponse struct{}"
	}

	return strings.Join([]string{"AddOrRemoveServicePermissionsResponse", string(data)}, " ")
}
