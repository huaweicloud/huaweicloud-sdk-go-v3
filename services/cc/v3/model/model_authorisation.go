package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Authorisation 授权
type Authorisation struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名字。
	Name string `json:"name"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 网络实例（VPC，VGW）的ID。
	InstanceId string `json:"instance_id"`

	// 实例所属项目ID。
	ProjectId string `json:"project_id"`

	// RegionID。
	RegionId string `json:"region_id"`

	// 实例创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 实例更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 实例所属账号ID。
	DomainId string `json:"domain_id"`

	// 云连接实例ID。
	CloudConnectionId string `json:"cloud_connection_id"`

	// 授权的状态。
	Status *string `json:"status,omitempty"`

	// 授权实例的类型。
	InstanceType *string `json:"instance_type,omitempty"`

	// 被授权云连接实例所属的账户ID。
	CloudConnectionDomainId *string `json:"cloud_connection_domain_id,omitempty"`

	// 是否已经被云连接加载。
	IsLoadedByCloudConnection *bool `json:"is_loaded_by_cloud_connection,omitempty"`
}

func (o Authorisation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Authorisation struct{}"
	}

	return strings.Join([]string{"Authorisation", string(data)}, " ")
}
