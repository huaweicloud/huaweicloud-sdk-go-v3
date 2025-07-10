package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddKeystorePermissionRequestBody 执行任务接口请求体
type AddKeystorePermissionRequestBody struct {

	// 是否有删除权限
	Delete bool `json:"delete"`

	// 文件密钥id
	KeystoreId string `json:"keystore_id"`

	// 是否有修改权限
	Modify bool `json:"modify"`

	// 是否有使用权限
	Usage bool `json:"usage"`

	// 用户名
	UserName string `json:"user_name"`

	// 是否有设置权限
	Setting bool `json:"setting"`

	// 是否可编辑
	CanAbsent bool `json:"can_absent"`
}

func (o AddKeystorePermissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddKeystorePermissionRequestBody struct{}"
	}

	return strings.Join([]string{"AddKeystorePermissionRequestBody", string(data)}, " ")
}
