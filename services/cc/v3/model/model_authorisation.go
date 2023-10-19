package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Authorisation 授权
type Authorisation struct {

	// 资源ID标识符。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 资源ID标识符。
	InstanceId string `json:"instance_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 实例所属帐号ID。
	DomainId string `json:"domain_id"`

	// 资源ID标识符。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 授权的状态。
	Status *string `json:"status,omitempty"`

	// 授权实例的类型。
	InstanceType *string `json:"instance_type,omitempty"`

	// 被授权云连接实例所属的账户ID。
	CloudConnectionDomainId *string `json:"cloud_connection_domain_id,omitempty"`
}

func (o Authorisation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authorisation struct{}"
	}

	return strings.Join([]string{"Authorisation", string(data)}, " ")
}
