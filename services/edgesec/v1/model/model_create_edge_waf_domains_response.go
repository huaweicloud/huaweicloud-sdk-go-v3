package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeWafDomainsResponse Response Object
type CreateEdgeWafDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateEdgeWafDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeWafDomainsResponse struct{}"
	}

	return strings.Join([]string{"CreateEdgeWafDomainsResponse", string(data)}, " ")
}
