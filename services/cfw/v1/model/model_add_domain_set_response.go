package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainSetResponse Response Object
type AddDomainSetResponse struct {
	Data           *DomainSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AddDomainSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainSetResponse struct{}"
	}

	return strings.Join([]string{"AddDomainSetResponse", string(data)}, " ")
}
