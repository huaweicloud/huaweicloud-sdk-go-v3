package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainResponse Response Object
type CreateDomainResponse struct {
	Domain         *CreateDomainResponseBodyContent `json:"domain,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainResponse", string(data)}, " ")
}
