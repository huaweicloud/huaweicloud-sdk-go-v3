package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertRequest Request Object
type CreateCertRequest struct {
	Body *CreateCertificateReq `json:"body,omitempty"`
}

func (o CreateCertRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertRequest struct{}"
	}

	return strings.Join([]string{"CreateCertRequest", string(data)}, " ")
}
