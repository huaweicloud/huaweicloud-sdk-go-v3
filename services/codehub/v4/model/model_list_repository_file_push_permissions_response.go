package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryFilePushPermissionsResponse Response Object
type ListRepositoryFilePushPermissionsResponse struct {

	// 仓库文件推送权限列表
	Body           *[]RepositoryFilePushPermissionDto `json:"body,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListRepositoryFilePushPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryFilePushPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryFilePushPermissionsResponse", string(data)}, " ")
}
