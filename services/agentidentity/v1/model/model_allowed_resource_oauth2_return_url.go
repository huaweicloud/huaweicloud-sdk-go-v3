package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowedResourceOauth2ReturnUrl The list item of allowed OAuth2 return URLs for resources associated with this workload identity.
type AllowedResourceOauth2ReturnUrl struct {
}

func (o AllowedResourceOauth2ReturnUrl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowedResourceOauth2ReturnUrl struct{}"
	}

	return strings.Join([]string{"AllowedResourceOauth2ReturnUrl", string(data)}, " ")
}
