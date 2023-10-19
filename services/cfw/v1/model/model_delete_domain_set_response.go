package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainSetResponse Response Object
type DeleteDomainSetResponse struct {
	Data           *DomainSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o DeleteDomainSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainSetResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainSetResponse", string(data)}, " ")
}
