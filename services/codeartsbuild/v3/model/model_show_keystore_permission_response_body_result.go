package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeystorePermissionResponseBodyResult 结果
type ShowKeystorePermissionResponseBodyResult struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 权限结集合
	PermissionList *[]AddKeystorePermissionResponseBody `json:"permission_list,omitempty"`
}

func (o ShowKeystorePermissionResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeystorePermissionResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ShowKeystorePermissionResponseBodyResult", string(data)}, " ")
}
