package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainTemplateResponse Response Object
type UpdateDomainTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainTemplateResponse", string(data)}, " ")
}
