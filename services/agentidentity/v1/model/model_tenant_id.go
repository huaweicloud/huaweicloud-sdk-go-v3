package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TenantId The tenant ID for the Microsoft OAuth2 provider.
type TenantId struct {
}

func (o TenantId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantId struct{}"
	}

	return strings.Join([]string{"TenantId", string(data)}, " ")
}
