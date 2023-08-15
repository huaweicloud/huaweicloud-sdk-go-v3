package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRepoDomainsResponse Response Object
type DeleteRepoDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRepoDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRepoDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteRepoDomainsResponse", string(data)}, " ")
}
