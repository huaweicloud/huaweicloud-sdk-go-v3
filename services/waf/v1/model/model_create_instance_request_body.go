package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceRequestBody CreateInstanceRequestBody
type CreateInstanceRequestBody struct {

	// 收费模式，当前仅支持按需收费（30）
	Chargemode *int32 `json:"chargemode,omitempty"`

	// 需要创建独享引擎的局点，例如：北京四（cn-north-4）
	Region string `json:"region"`

	// 需要创建独享引擎的可用区，例如：北京四可用区1（cn-north-4a）
	AvailableZone string `json:"available_zone"`

	// 独享引擎CPU架构，例如：x86与arm
	Arch string `json:"arch"`

	// 独享引擎名称前缀
	Instancename string `json:"instancename"`

	// 独享引擎版本规格   - 企业版：waf.instance.enterprise   - 专业版：waf.instance.professional
	Specification string `json:"specification"`

	// 独享引擎ECS规格，实例规格企业版对应8U16G的ecs规格，专业版对应2U4G的ecs规格（通过调用ECS的ListFlavors接口获取应8U16G的ecs和2U4G的ecs对应规格id）
	CpuFlavor string `json:"cpu_flavor"`

	// 独享引擎所在VPC的ID（通过调用虚拟私有云ListVpcs接口获取所有的VPC列表查询VPC的ID，如果需要关联企业项目，则调用虚拟私有云的接口也需要关联企业项目ID）
	VpcId string `json:"vpc_id"`

	// 独享引擎所在VPC内的子网ID（通过调用虚拟私有云ListSubnets接口获取所有的子网列表查询子网的ID，如果需要关联企业项目，则调用虚拟私有云的接口也需要关联企业项目ID）
	SubnetId string `json:"subnet_id"`

	// 独享引擎需要绑定的安全组ID（通过调用虚拟私有云ListSecurityGroups接口获取所有的安全组列表查询安全组的ID，如果需要关联企业项目，则调用虚拟私有云的接口也需要关联企业项目ID）
	SecurityGroup []string `json:"security_group"`

	// 申请的独享引擎数量
	Count int32 `json:"count"`

	// 是否为资源租户类，默认值为false。   - true: 资源租户类   - false: 普通租户类
	ResTenant *bool `json:"res_tenant,omitempty"`
}

func (o CreateInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceRequestBody", string(data)}, " ")
}
