package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDomainsResponse Response Object
type AddDomainsResponse struct {
	Data           *DomainSetResponseData `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o AddDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDomainsResponse struct{}"
	}

	return strings.Join([]string{"AddDomainsResponse", string(data)}, " ")
}
