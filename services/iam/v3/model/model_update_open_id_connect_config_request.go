package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateOpenIdConnectConfigRequest struct {

	// 身份提供商ID
	IdpId string `json:"idp_id" xml:"idp_id"`

	Body *UpdateOpenIdConnectConfigRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateOpenIdConnectConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOpenIdConnectConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateOpenIdConnectConfigRequest", string(data)}, " ")
}
