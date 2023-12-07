package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityPermissionSetsResponse Response Object
type ListSecurityPermissionSetsResponse struct {

	// 总条数
	Total *int32 `json:"total,omitempty"`

	// 权限集列表
	PermissionSets *[]PermissionSet `json:"permission_sets,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListSecurityPermissionSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPermissionSetsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPermissionSetsResponse", string(data)}, " ")
}
