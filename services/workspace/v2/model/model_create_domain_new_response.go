package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainNewResponse Response Object
type CreateDomainNewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDomainNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainNewResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainNewResponse", string(data)}, " ")
}
