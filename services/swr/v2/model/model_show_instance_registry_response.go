package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRegistryResponse Response Object
type ShowInstanceRegistryResponse struct {

	// 仓库ID
	Id *int32 `json:"id,omitempty"`

	// 仓库名称
	Name *string `json:"name,omitempty"`

	// 仓库描述
	Description *string `json:"description,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 仓库类型
	Type *string `json:"type,omitempty"`

	// 仓库地址
	Url *string `json:"url,omitempty"`

	Credential *Credential `json:"credential,omitempty"`

	// 是否验证远程证书
	Insecure *bool `json:"insecure,omitempty"`

	// 仓库状态，healthy或unhealthy
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRegistryResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceRegistryResponse", string(data)}, " ")
}
