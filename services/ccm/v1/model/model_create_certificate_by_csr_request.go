package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateByCsrRequest Request Object
type CreateCertificateByCsrRequest struct {
	Body *CreateCertificateByCsrRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateByCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateByCsrRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateByCsrRequest", string(data)}, " ")
}
