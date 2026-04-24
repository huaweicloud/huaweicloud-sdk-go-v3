package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDNatResponse Response Object
type BindDNatResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BindDNatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDNatResponse struct{}"
	}

	return strings.Join([]string{"BindDNatResponse", string(data)}, " ")
}
