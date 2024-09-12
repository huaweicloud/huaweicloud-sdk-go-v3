package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainsResponse Response Object
type DeleteDomainsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainsResponse", string(data)}, " ")
}
