package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityResourcePermissionsResponse Response Object
type ListSecurityResourcePermissionsResponse struct {

	// 空间资源权限策略总条数
	Total *int32 `json:"total,omitempty"`

	// 空间资源权限策略列表
	Policies       *[]PermissionResourcePolicy `json:"policies,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListSecurityResourcePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityResourcePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityResourcePermissionsResponse", string(data)}, " ")
}
