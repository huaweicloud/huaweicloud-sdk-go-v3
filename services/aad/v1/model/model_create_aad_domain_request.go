package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAadDomainRequest Request Object
type CreateAadDomainRequest struct {
	Body *HostBody `json:"body,omitempty"`
}

func (o CreateAadDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAadDomainRequest struct{}"
	}

	return strings.Join([]string{"CreateAadDomainRequest", string(data)}, " ")
}
