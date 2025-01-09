package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantProfilesResponse Response Object
type ListTenantProfilesResponse struct {

	// 查询功能开关列表。
	Body           map[string]bool `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTenantProfilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantProfilesResponse struct{}"
	}

	return strings.Join([]string{"ListTenantProfilesResponse", string(data)}, " ")
}
