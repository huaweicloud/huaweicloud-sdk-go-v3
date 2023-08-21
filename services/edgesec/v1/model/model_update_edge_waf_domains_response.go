package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeWafDomainsResponse Response Object
type UpdateEdgeWafDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEdgeWafDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeWafDomainsResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeWafDomainsResponse", string(data)}, " ")
}
