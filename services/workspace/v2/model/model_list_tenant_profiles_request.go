package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantProfilesRequest Request Object
type ListTenantProfilesRequest struct {
}

func (o ListTenantProfilesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantProfilesRequest struct{}"
	}

	return strings.Join([]string{"ListTenantProfilesRequest", string(data)}, " ")
}
