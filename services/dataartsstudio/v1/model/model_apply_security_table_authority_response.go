package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplySecurityTableAuthorityResponse Response Object
type ApplySecurityTableAuthorityResponse struct {

	// 描述
	Describe *string `json:"describe,omitempty"`

	// 审批页面地址
	PermissionCenterUrl *string `json:"permission_center_url,omitempty"`

	// 工作空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 工单id
	ApplicationId  *string `json:"application_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplySecurityTableAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplySecurityTableAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ApplySecurityTableAuthorityResponse", string(data)}, " ")
}
