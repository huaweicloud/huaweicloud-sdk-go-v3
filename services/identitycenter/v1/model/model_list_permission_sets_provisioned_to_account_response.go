package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionSetsProvisionedToAccountResponse Response Object
type ListPermissionSetsProvisionedToAccountResponse struct {

	// 权限集列表
	PermissionSets *[]string `json:"permission_sets,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPermissionSetsProvisionedToAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetsProvisionedToAccountResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionSetsProvisionedToAccountResponse", string(data)}, " ")
}
