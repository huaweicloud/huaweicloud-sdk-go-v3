package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyV5Request Request Object
type CreateAgencyV5Request struct {
	Body *CreateAgencyReqBody `json:"body,omitempty"`
}

func (o CreateAgencyV5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyV5Request struct{}"
	}

	return strings.Join([]string{"CreateAgencyV5Request", string(data)}, " ")
}
