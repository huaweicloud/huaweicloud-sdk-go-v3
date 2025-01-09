package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantProfileRequest Request Object
type ListTenantProfileRequest struct {
}

func (o ListTenantProfileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantProfileRequest struct{}"
	}

	return strings.Join([]string{"ListTenantProfileRequest", string(data)}, " ")
}
