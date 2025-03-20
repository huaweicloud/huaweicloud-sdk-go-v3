package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateServiceLinkedAgencyV5Request Request Object
type CreateServiceLinkedAgencyV5Request struct {
	Body *CreateServiceLinkedAgencyReqBody `json:"body,omitempty"`
}

func (o CreateServiceLinkedAgencyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateServiceLinkedAgencyV5Request struct{}"
	}

	return strings.Join([]string{"CreateServiceLinkedAgencyV5Request", string(data)}, " ")
}
