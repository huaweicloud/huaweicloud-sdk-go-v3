package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializeTenantResponse Response Object
type InitializeTenantResponse struct {

	// 租户ID 同tenant_id。
	ProjectId *string `json:"project_id,omitempty"`

	// 企业是否激活：active(激活)，inactive(未激活)。
	ServiceStatus  *string `json:"service_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InitializeTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializeTenantResponse struct{}"
	}

	return strings.Join([]string{"InitializeTenantResponse", string(data)}, " ")
}
