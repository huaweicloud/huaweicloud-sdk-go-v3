package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {

	// 委托ID。
	AgencyId *string `json:"agency_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
