package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainTemplateResponse Response Object
type DeleteDomainTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainTemplateResponse", string(data)}, " ")
}
