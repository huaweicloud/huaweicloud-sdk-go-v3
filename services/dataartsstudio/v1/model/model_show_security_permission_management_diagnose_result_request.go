package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityPermissionManagementDiagnoseResultRequest Request Object
type ShowSecurityPermissionManagementDiagnoseResultRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowSecurityPermissionManagementDiagnoseResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPermissionManagementDiagnoseResultRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityPermissionManagementDiagnoseResultRequest", string(data)}, " ")
}
