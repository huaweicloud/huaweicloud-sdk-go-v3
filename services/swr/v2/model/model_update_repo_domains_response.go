package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepoDomainsResponse Response Object
type UpdateRepoDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateRepoDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoDomainsResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepoDomainsResponse", string(data)}, " ")
}
