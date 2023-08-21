package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeDDoSDomainsResponse Response Object
type UpdateEdgeDDoSDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateEdgeDDoSDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeDDoSDomainsResponse struct{}"
	}

	return strings.Join([]string{"UpdateEdgeDDoSDomainsResponse", string(data)}, " ")
}
