package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnterpriseRealnameAuthenticationRequest struct {
	Body *ApplyEnterpriseRealnameAuthsReq `json:"body,omitempty"`
}

func (o CreateEnterpriseRealnameAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRealnameAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRealnameAuthenticationRequest", string(data)}, " ")
}
