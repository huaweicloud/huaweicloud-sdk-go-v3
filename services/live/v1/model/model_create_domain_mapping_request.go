package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainMappingRequest Request Object
type CreateDomainMappingRequest struct {
	Body *DomainMapping `json:"body,omitempty"`
}

func (o CreateDomainMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainMappingRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainMappingRequest", string(data)}, " ")
}
