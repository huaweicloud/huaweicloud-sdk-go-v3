package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAgencyCredentialsResponse Response Object
type GetAgencyCredentialsResponse struct {
	AgencyCredentials *AgencyCredentials `json:"agency_credentials,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o GetAgencyCredentialsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAgencyCredentialsResponse struct{}"
	}

	return strings.Join([]string{"GetAgencyCredentialsResponse", string(data)}, " ")
}
