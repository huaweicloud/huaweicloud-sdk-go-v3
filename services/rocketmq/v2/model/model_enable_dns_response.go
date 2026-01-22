package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDnsResponse Response Object
type EnableDnsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDnsResponse struct{}"
	}

	return strings.Join([]string{"EnableDnsResponse", string(data)}, " ")
}
