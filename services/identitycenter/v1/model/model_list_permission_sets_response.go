package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionSetsResponse Response Object
type ListPermissionSetsResponse struct {

	// 权限集列表
	PermissionSets *[]PermissionSetDto `json:"permission_sets,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPermissionSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetsResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionSetsResponse", string(data)}, " ")
}
