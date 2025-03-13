package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySecurityAdminRequest Request Object
type ModifySecurityAdminRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *DlsAdmin `json:"body,omitempty"`
}

func (o ModifySecurityAdminRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityAdminRequest struct{}"
	}

	return strings.Join([]string{"ModifySecurityAdminRequest", string(data)}, " ")
}
