package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyV5Response Response Object
type CreateAgencyV5Response struct {
	Agency         *Agency `json:"agency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyV5Response struct{}"
	}

	return strings.Join([]string{"CreateAgencyV5Response", string(data)}, " ")
}
