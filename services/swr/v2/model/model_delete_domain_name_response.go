package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainNameResponse Response Object
type DeleteDomainNameResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDomainNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainNameResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainNameResponse", string(data)}, " ")
}
