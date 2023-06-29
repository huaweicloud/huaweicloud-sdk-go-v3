package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoRenewalResourcesResponse Response Object
type AutoRenewalResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AutoRenewalResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoRenewalResourcesResponse struct{}"
	}

	return strings.Join([]string{"AutoRenewalResourcesResponse", string(data)}, " ")
}
