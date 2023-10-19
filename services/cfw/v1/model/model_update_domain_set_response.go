package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainSetResponse Response Object
type UpdateDomainSetResponse struct {
	Data           *DomainSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateDomainSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainSetResponse", string(data)}, " ")
}
