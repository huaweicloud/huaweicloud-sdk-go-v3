package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {

	// OBS当前的授权结果。 - **true** - **false**
	AgencyGranted  *string `json:"agency_granted,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
