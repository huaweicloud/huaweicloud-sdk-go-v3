package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterDnsResponse Response Object
type CreateClusterDnsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateClusterDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterDnsResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterDnsResponse", string(data)}, " ")
}
