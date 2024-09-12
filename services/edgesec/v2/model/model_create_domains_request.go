package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainsRequest Request Object
type CreateDomainsRequest struct {
	Body *CreateDomainRequestBody `json:"body,omitempty"`
}

func (o CreateDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsRequest struct{}"
	}

	return strings.Join([]string{"CreateDomainsRequest", string(data)}, " ")
}
