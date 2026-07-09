package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CsbResource struct {

	// 资源标签
	Tags *[]CsbResourceTag `json:"tags,omitempty"`

	// 防护资源对象列表
	AffectedResources *[]AffectedResource `json:"affected_resources,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 资源描述
	Description *string `json:"description,omitempty"`

	Environment *ResourceEnvironment `json:"environment,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 所属云服务，dbss
	Provider *string `json:"provider,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 安全组ID
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 资源类型   - cloudservers: 审计   - dbEncrypt: 加密   - dbOm: 运维
	Type *string `json:"type,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 资源URN
	Urn *string `json:"urn,omitempty"`

	// 资源URN扩展
	Urnext *string `json:"urnext,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o CsbResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CsbResource struct{}"
	}

	return strings.Join([]string{"CsbResource", string(data)}, " ")
}
