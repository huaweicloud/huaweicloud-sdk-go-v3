package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnterpriseRealnameAuthenticationRequest struct {
	Body *ApplyEnterpriseRealnameAuthsReq `json:"body,omitempty" xml:"body"`
}

func (o CreateEnterpriseRealnameAuthenticationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseRealnameAuthenticationRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseRealnameAuthenticationRequest", string(data)}, " ")
}
