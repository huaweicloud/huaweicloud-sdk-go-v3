package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreSearchResponseBodyResultPermission 权限
type ListKeystoreSearchResponseBodyResultPermission struct {

	// 文件ID
	KeystoreId *string `json:"keystore_id,omitempty"`

	// 是否有设置权限
	Setting *bool `json:"setting,omitempty"`

	// 是否有删除权限
	Delete *bool `json:"delete,omitempty"`

	// 是否有修改权限
	Modify *bool `json:"modify,omitempty"`

	// 是否有使用权限
	Usage *bool `json:"usage,omitempty"`
}

func (o ListKeystoreSearchResponseBodyResultPermission) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreSearchResponseBodyResultPermission struct{}"
	}

	return strings.Join([]string{"ListKeystoreSearchResponseBodyResultPermission", string(data)}, " ")
}
