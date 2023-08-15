package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityObsAgencyResponse Response Object
type ShowCertificateAuthorityObsAgencyResponse struct {

	// OBS当前的授权结果。 - **true** - **false**
	AgencyGranted  *string `json:"agency_granted,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertificateAuthorityObsAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityObsAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityObsAgencyResponse", string(data)}, " ")
}
