package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityObsAgencyRequest Request Object
type CreateCertificateAuthorityObsAgencyRequest struct {
}

func (o CreateCertificateAuthorityObsAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityObsAgencyRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityObsAgencyRequest", string(data)}, " ")
}
