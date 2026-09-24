package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInternalEndpointPermissionsResponse Response Object
type ListInternalEndpointPermissionsResponse struct {

	// 权限列表
	Permissions *[]PermissionItem `json:"permissions,omitempty"`

	// 满足查询条件的白名单总条数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInternalEndpointPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInternalEndpointPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListInternalEndpointPermissionsResponse", string(data)}, " ")
}
