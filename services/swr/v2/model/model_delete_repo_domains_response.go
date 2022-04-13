package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
