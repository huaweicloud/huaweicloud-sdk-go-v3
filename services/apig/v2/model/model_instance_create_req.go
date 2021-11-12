package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceCreateReq struct {
	// 实例描述

	Description *string `json:"description,omitempty"`
	// '维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。'

	MaintainBegin *string `json:"maintain_begin,omitempty"`
	// '维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。'  '在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次'。

	MaintainEnd *string `json:"maintain_end,omitempty"`
	// 实例名称

	InstanceName *string `json:"instance_name,omitempty"`
	// 实例编号，不填写自动生成

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例规格： - BASIC：基础版实例 - PROFESSIONAL：专业版实例 - ENTERPRISE：企业版实例 - PLATINUM：铂金版实例 - BASIC_IPV6：基础版IPV6实例 - PROFESSIONAL_IPV6：专业版IPV6实例 - ENTERPRISE_IPV6：企业版IPV6实例 - PLATINUM_IPV6：铂金版IPV6实例

	SpecId *InstanceCreateReqSpecId `json:"spec_id,omitempty"`
	// 虚拟私有云ID。  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询VPC列表”章节。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网的网络ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询子网列表”章节。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 指定实例所属的安全组。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询安全组列表”章节。

	SecurityGroupId *string `json:"security_group_id,omitempty"`
	// 弹性公网IP ID。  实例需要开启公网访问时需要填写，绑定后使用者可以通过该入口从公网访问APIG实例中的API等资源  获取方法：登录虚拟私有云服务的控制台界面，在弹性公网IP的详情页面查找弹性公网IP ID。

	EipId *string `json:"eip_id,omitempty"`
	// 企业项目ID，企业帐号必填。  获取方法如下： - 方法1：登录企业项目管理界面，在项目管理详情页面查找项目ID。 - 方法2：通过企业项目管理的API接口查询，具体方法请参见《企业管理API参考》的“查询企业项目列表”章节。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 可用区列表。  可用区指在同一地域下，电力、网络隔离的物理区域，可用区之内内网互通，不同可用区之间物理隔离。选择多个AZ部署可以有效提升可靠性。  获取方法：通过文档中实例管理的可用区列表接口查询。

	AvailableZoneIds *[]string `json:"available_zone_ids,omitempty"`
	// 出公网带宽  实例需要开启出公网功能时需要填写，绑定后使用者可以利用该出口访问公网上的互联网资源

	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`
	// 公网访问是否支持IPv6。  当前仅部分region部分可用区支持IPv6

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
}

func (o InstanceCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCreateReq struct{}"
	}

	return strings.Join([]string{"InstanceCreateReq", string(data)}, " ")
}

type InstanceCreateReqSpecId struct {
	value string
}

type InstanceCreateReqSpecIdEnum struct {
	BASIC             InstanceCreateReqSpecId
	PROFESSIONAL      InstanceCreateReqSpecId
	ENTERPRISE        InstanceCreateReqSpecId
	PLATINUM          InstanceCreateReqSpecId
	BASIC_IPV6        InstanceCreateReqSpecId
	PROFESSIONAL_IPV6 InstanceCreateReqSpecId
	ENTERPRISE_IPV6   InstanceCreateReqSpecId
	PLATINUM_IPV6     InstanceCreateReqSpecId
}

func GetInstanceCreateReqSpecIdEnum() InstanceCreateReqSpecIdEnum {
	return InstanceCreateReqSpecIdEnum{
		BASIC: InstanceCreateReqSpecId{
			value: "BASIC",
		},
		PROFESSIONAL: InstanceCreateReqSpecId{
			value: "PROFESSIONAL",
		},
		ENTERPRISE: InstanceCreateReqSpecId{
			value: "ENTERPRISE",
		},
		PLATINUM: InstanceCreateReqSpecId{
			value: "PLATINUM",
		},
		BASIC_IPV6: InstanceCreateReqSpecId{
			value: "BASIC_IPV6",
		},
		PROFESSIONAL_IPV6: InstanceCreateReqSpecId{
			value: "PROFESSIONAL_IPV6",
		},
		ENTERPRISE_IPV6: InstanceCreateReqSpecId{
			value: "ENTERPRISE_IPV6",
		},
		PLATINUM_IPV6: InstanceCreateReqSpecId{
			value: "PLATINUM_IPV6",
		},
	}
}

func (c InstanceCreateReqSpecId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceCreateReqSpecId) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
