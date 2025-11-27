package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainByDuplicateRequest Request Object
type CreateDomainByDuplicateRequest struct {
	Body *DuplicateDomainRequestBody `json:"body,omitempty"`
}

func (o CreateDomainByDuplicateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainByDuplicateRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainByDuplicateRequest", string(data)}, " ")
}
