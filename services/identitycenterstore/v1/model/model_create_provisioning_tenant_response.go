package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProvisioningTenantResponse Response Object
type CreateProvisioningTenantResponse struct {

	// 创建时间
	CreationTime float32 `json:"creation_time,omitempty"`

	// SCIM终端节点
	ScimEndpoint *string `json:"scim_endpoint,omitempty"`

	// 开启自动预配后生成的全局唯一标识符（ID）
	TenantId       *string `json:"tenant_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProvisioningTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProvisioningTenantResponse struct{}"
	}

	return strings.Join([]string{"CreateProvisioningTenantResponse", string(data)}, " ")
}
