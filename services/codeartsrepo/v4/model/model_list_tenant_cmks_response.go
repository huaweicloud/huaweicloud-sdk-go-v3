package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantCmksResponse Response Object
type ListTenantCmksResponse struct {
	Body           *[]TenantCmkDto `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTenantCmksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantCmksResponse struct{}"
	}

	return strings.Join([]string{"ListTenantCmksResponse", string(data)}, " ")
}
