package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionSetProvisioningStatusResponse Response Object
type ListPermissionSetProvisioningStatusResponse struct {

	// 权限集授权状态
	PermissionSetsProvisioningStatus *[]PermissionSetProvisioningStatusMetadataDto `json:"permission_sets_provisioning_status,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListPermissionSetProvisioningStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionSetProvisioningStatusResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionSetProvisioningStatusResponse", string(data)}, " ")
}
