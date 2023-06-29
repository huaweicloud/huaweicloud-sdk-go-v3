package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityObsAgencyResponse Response Object
type CreateCertificateAuthorityObsAgencyResponse struct {

	// 创建OBS委托，由IAM返回的授权ID。
	AgencyId       *string `json:"agency_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCertificateAuthorityObsAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityObsAgencyResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityObsAgencyResponse", string(data)}, " ")
}
