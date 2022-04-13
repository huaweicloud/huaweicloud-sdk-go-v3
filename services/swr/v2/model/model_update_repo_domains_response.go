package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
