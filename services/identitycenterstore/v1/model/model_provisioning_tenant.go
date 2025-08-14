package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProvisioningTenant struct {

	// 创建时间
	CreationTime float32 `json:"creation_time"`

	// SCIM终端节点
	ScimEndpoint string `json:"scim_endpoint"`

	// 开启自动预配后生成的全局唯一标识符（ID）
	TenantId string `json:"tenant_id"`
}

func (o ProvisioningTenant) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProvisioningTenant struct{}"
	}

	return strings.Join([]string{"ProvisioningTenant", string(data)}, " ")
}
