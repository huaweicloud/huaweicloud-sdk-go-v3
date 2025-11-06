package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateRepositoryFilePushPermissionsResponse Response Object
type BatchUpdateRepositoryFilePushPermissionsResponse struct {

	// 仓库文件推送权限列表
	Body           *[]RepositoryFilePushPermissionDto `json:"body,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o BatchUpdateRepositoryFilePushPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRepositoryFilePushPermissionsResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateRepositoryFilePushPermissionsResponse", string(data)}, " ")
}
