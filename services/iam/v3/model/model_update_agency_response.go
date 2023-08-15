package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyResponse Response Object
type UpdateAgencyResponse struct {
	Agency         *AgencyResult `json:"agency,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o UpdateAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyResponse", string(data)}, " ")
}
