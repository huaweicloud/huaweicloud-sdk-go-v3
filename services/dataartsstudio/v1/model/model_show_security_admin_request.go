package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityAdminRequest Request Object
type ShowSecurityAdminRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o ShowSecurityAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityAdminRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityAdminRequest", string(data)}, " ")
}
