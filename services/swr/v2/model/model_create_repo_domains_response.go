package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRepoDomainsResponse Response Object
type CreateRepoDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateRepoDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoDomainsResponse struct{}"
	}

	return strings.Join([]string{"CreateRepoDomainsResponse", string(data)}, " ")
}
