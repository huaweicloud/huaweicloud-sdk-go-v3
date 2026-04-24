package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainNewRequest Request Object
type CreateDomainNewRequest struct {
	Body *CreateDomainNewReq `json:"body,omitempty"`
}

func (o CreateDomainNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainNewRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainNewRequest", string(data)}, " ")
}
