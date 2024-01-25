package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCertForDomainResponse Response Object
type SetCertForDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetCertForDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCertForDomainResponse struct{}"
	}

	return strings.Join([]string{"SetCertForDomainResponse", string(data)}, " ")
}
