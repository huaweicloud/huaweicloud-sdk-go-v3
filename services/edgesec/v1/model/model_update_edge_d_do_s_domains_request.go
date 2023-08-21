package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeDDoSDomainsRequest Request Object
type UpdateEdgeDDoSDomainsRequest struct {

	// 域名ID
	Domainid string `json:"domainid"`

	Body *UpdateEdgeDDoSDomainsRequestBody `json:"body,omitempty"`
}

func (o UpdateEdgeDDoSDomainsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeDDoSDomainsRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeDDoSDomainsRequest", string(data)}, " ")
}
