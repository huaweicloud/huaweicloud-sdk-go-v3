package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainsResponse Response Object
type CreateDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainsResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainsResponse", string(data)}, " ")
}
