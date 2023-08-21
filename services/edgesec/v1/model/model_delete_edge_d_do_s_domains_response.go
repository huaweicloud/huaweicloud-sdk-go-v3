package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEdgeDDoSDomainsResponse Response Object
type DeleteEdgeDDoSDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEdgeDDoSDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEdgeDDoSDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeDDoSDomainsResponse", string(data)}, " ")
}
