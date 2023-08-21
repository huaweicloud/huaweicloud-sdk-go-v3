package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeWafDomainsRequest Request Object
type UpdateEdgeWafDomainsRequest struct {

	// 域名
	Domainid string `json:"domainid"`

	Body *UpdateEdgeWafDomainsRequestBody `json:"body,omitempty"`
}

func (o UpdateEdgeWafDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeWafDomainsRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeWafDomainsRequest", string(data)}, " ")
}
