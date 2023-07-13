package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeEnterpriseRealnameAuthenticationRequest Request Object
type ChangeEnterpriseRealnameAuthenticationRequest struct {
	Body *ChangeEnterpriseRealnameAuthsReq `json:"body,omitempty"`
}

func (o ChangeEnterpriseRealnameAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeEnterpriseRealnameAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"ChangeEnterpriseRealnameAuthenticationRequest", string(data)}, " ")
}
