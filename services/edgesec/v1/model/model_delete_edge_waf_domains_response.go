package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeWafDomainsResponse Response Object
type DeleteEdgeWafDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeWafDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeWafDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeWafDomainsResponse", string(data)}, " ")
}
