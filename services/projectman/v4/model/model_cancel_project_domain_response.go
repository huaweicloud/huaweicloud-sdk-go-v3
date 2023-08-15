package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelProjectDomainResponse Response Object
type CancelProjectDomainResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelProjectDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelProjectDomainResponse struct{}"
	}

	return strings.Join([]string{"CancelProjectDomainResponse", string(data)}, " ")
}
