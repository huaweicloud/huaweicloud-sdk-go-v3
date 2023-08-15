package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClusterDnsResponse Response Object
type UpdateClusterDnsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateClusterDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterDnsResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterDnsResponse", string(data)}, " ")
}
