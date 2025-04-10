package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityMemberPermissionExpireTimeRequest Request Object
type UpdateSecurityMemberPermissionExpireTimeRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BatchUpdatePermissionExpireTimeDto `json:"body,omitempty"`
}

func (o UpdateSecurityMemberPermissionExpireTimeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityMemberPermissionExpireTimeRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityMemberPermissionExpireTimeRequest", string(data)}, " ")
}
