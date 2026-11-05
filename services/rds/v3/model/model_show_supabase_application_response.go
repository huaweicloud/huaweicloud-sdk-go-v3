package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSupabaseApplicationResponse Response Object
type ShowSupabaseApplicationResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	// 运行状态。
	Status *string `json:"status,omitempty"`

	EnableSsl *bool `json:"enable_ssl,omitempty"`

	CertId *string `json:"cert_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	StorageComponent *SupabaseStorageComponentInfo `json:"storage_component,omitempty"`

	ManagedInstance *SupabaseManagedInstanceInfo `json:"managed_instance,omitempty"`

	UpdatedTime    *int64 `json:"updated_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSupabaseApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSupabaseApplicationResponse struct{}"
	}

	return strings.Join([]string{"ShowSupabaseApplicationResponse", string(data)}, " ")
}
