package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeDDoSDomainsResponse Response Object
type CreateEdgeDDoSDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateEdgeDDoSDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeDDoSDomainsResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeDDoSDomainsResponse", string(data)}, " ")
}
