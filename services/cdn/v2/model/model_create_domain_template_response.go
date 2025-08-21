package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainTemplateResponse Response Object
type CreateDomainTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDomainTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainTemplateResponse", string(data)}, " ")
}
