package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgreementRequest Request Object
type CreateAgreementRequest struct {
	Body *TenantAgreementBody `json:"body,omitempty"`
}

func (o CreateAgreementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgreementRequest struct{}"
	}

	return strings.Join([]string{"CreateAgreementRequest", string(data)}, " ")
}
