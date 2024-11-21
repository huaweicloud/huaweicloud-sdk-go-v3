package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyResponse Response Object
type ShowAgencyResponse struct {
	Agency         *AgencyShowResult `json:"agency,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyResponse struct{}"
	}

	return strings.Join([]string{"ShowAgencyResponse", string(data)}, " ")
}
