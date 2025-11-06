package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryResourcePermissionsResponse Response Object
type ListRepositoryResourcePermissionsResponse struct {

	// 仓库权限矩阵配置列表
	Body           *[]ResponsePermissionInfo `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListRepositoryResourcePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryResourcePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryResourcePermissionsResponse", string(data)}, " ")
}
