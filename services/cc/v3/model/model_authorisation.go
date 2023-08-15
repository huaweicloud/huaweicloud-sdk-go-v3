package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Authorisation 授权
type Authorisation struct {

	// 授权的ID。
	Id *string `json:"id,omitempty"`

	// 授权的名称。
	Name *string `json:"name,omitempty"`

	// 授权的描述信息。
	Description *string `json:"description,omitempty"`

	// 授权的状态。
	Status *string `json:"status,omitempty"`

	// 创建授权的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新授权的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 授权者的账户ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 授权实例的ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 授权实例的类型。
	InstanceType *string `json:"instance_type,omitempty"`

	// 授权实例所属Region。
	RegionId *string `json:"region_id,omitempty"`

	// 授权实例所属项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 被授权云连接实例所属的账户ID。
	CloudConnectionDomainId *string `json:"cloud_connection_domain_id,omitempty"`

	// 被授权云连接实例ID。
	CloudConnectionId *string `json:"cloud_connection_id,omitempty"`
}

func (o Authorisation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authorisation struct{}"
	}

	return strings.Join([]string{"Authorisation", string(data)}, " ")
}
