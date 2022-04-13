package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
