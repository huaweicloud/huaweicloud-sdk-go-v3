package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddKeystorePermissionResponseBody 执行任务接口请求体
type AddKeystorePermissionResponseBody struct {

	// 权限id
	Id *string `json:"id,omitempty"`

	// 编辑权限
	Setting *bool `json:"setting,omitempty"`

	// 是否有删除权限
	Delete *bool `json:"delete,omitempty"`

	// 文件密钥id
	KeystoreId *string `json:"keystore_id,omitempty"`

	// 是否有修改权限
	Modify *bool `json:"modify,omitempty"`

	// 是否有使用权限
	Usage *bool `json:"usage,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 是否是创建者
	IsCreator *bool `json:"is_creator,omitempty"`

	// 是否可编辑
	CanAbsent *bool `json:"can_absent,omitempty"`
}

func (o AddKeystorePermissionResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddKeystorePermissionResponseBody struct{}"
	}

	return strings.Join([]string{"AddKeystorePermissionResponseBody", string(data)}, " ")
}
