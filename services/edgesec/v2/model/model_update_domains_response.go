package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainsResponse Response Object
type UpdateDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainsResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainsResponse", string(data)}, " ")
}
