package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceLinkedAgencyV5Response Response Object
type CreateServiceLinkedAgencyV5Response struct {
	Agency         *Agency `json:"agency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateServiceLinkedAgencyV5Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceLinkedAgencyV5Response struct{}"
	}

	return strings.Join([]string{"CreateServiceLinkedAgencyV5Response", string(data)}, " ")
}
