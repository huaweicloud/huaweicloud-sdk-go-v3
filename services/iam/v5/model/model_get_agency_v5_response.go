package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAgencyV5Response Response Object
type GetAgencyV5Response struct {
	Agency         *AgencyEx `json:"agency,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o GetAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAgencyV5Response struct{}"
	}

	return strings.Join([]string{"GetAgencyV5Response", string(data)}, " ")
}
