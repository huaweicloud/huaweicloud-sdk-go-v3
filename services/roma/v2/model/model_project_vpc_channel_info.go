package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ProjectVpcChannelInfo struct {

	// VPC通道的名称。  长度为3 ~ 64位的字符串，字符串由中文、英文字母、数字、中划线、下划线组成，且只能以英文或中文开头。 > 中文字符必须为UTF-8或者unicode编码。
	Name string `json:"name"`

	// VPC通道中主机的端口号。  取值范围1 ~ 65535。
	Port int32 `json:"port"`

	// 分发算法。 - 1：加权轮询（wrr） - 2：加权最少连接（wleastconn） - 3：源地址哈希（source） - 4：URI哈希（uri）
	BalanceStrategy ProjectVpcChannelInfoBalanceStrategy `json:"balance_strategy"`

	// VPC通道的成员类型。[site场景必须修改成IP类型](tag:Site) - ip - ecs
	MemberType ProjectVpcChannelInfoMemberType `json:"member_type"`

	// VPC通道的字典编码  支持英文，数字，特殊字符（-_.）  暂不支持
	DictCode *string `json:"dict_code,omitempty"`

	// VPC通道的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// VPC通道的编号
	Id *string `json:"id,omitempty"`

	// VPC通道的状态。 - 1：正常 - 2：异常
	Status *ProjectVpcChannelInfoStatus `json:"status,omitempty"`

	// 后端云服务器组列表。
	MemberGroups *[]MemberGroupInfo `json:"member_groups,omitempty"`

	// 实例编号
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 后端实例列表。
	Members *[]VpcMemberInfo `json:"members,omitempty"`

	VpcHealthConfig *VpcHealthConfigInfo `json:"vpc_health_config,omitempty"`

	MicroserviceInfo *MicroServiceInfo `json:"microservice_info,omitempty"`

	// vpc通道类型。 - BUILTIN：BUILTIN通道类型 - MICROSERVICE：微服务类型
	Type *ProjectVpcChannelInfoType `json:"type,omitempty"`
}

func (o ProjectVpcChannelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectVpcChannelInfo struct{}"
	}

	return strings.Join([]string{"ProjectVpcChannelInfo", string(data)}, " ")
}

type ProjectVpcChannelInfoBalanceStrategy struct {
	value int32
}

type ProjectVpcChannelInfoBalanceStrategyEnum struct {
	E_1 ProjectVpcChannelInfoBalanceStrategy
	E_2 ProjectVpcChannelInfoBalanceStrategy
	E_3 ProjectVpcChannelInfoBalanceStrategy
	E_4 ProjectVpcChannelInfoBalanceStrategy
}

func GetProjectVpcChannelInfoBalanceStrategyEnum() ProjectVpcChannelInfoBalanceStrategyEnum {
	return ProjectVpcChannelInfoBalanceStrategyEnum{
		E_1: ProjectVpcChannelInfoBalanceStrategy{
			value: 1,
		}, E_2: ProjectVpcChannelInfoBalanceStrategy{
			value: 2,
		}, E_3: ProjectVpcChannelInfoBalanceStrategy{
			value: 3,
		}, E_4: ProjectVpcChannelInfoBalanceStrategy{
			value: 4,
		},
	}
}

func (c ProjectVpcChannelInfoBalanceStrategy) Value() int32 {
	return c.value
}

func (c ProjectVpcChannelInfoBalanceStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcChannelInfoBalanceStrategy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ProjectVpcChannelInfoMemberType struct {
	value string
}

type ProjectVpcChannelInfoMemberTypeEnum struct {
	IP  ProjectVpcChannelInfoMemberType
	ECS ProjectVpcChannelInfoMemberType
}

func GetProjectVpcChannelInfoMemberTypeEnum() ProjectVpcChannelInfoMemberTypeEnum {
	return ProjectVpcChannelInfoMemberTypeEnum{
		IP: ProjectVpcChannelInfoMemberType{
			value: "ip",
		},
		ECS: ProjectVpcChannelInfoMemberType{
			value: "ecs",
		},
	}
}

func (c ProjectVpcChannelInfoMemberType) Value() string {
	return c.value
}

func (c ProjectVpcChannelInfoMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcChannelInfoMemberType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ProjectVpcChannelInfoStatus struct {
	value int32
}

type ProjectVpcChannelInfoStatusEnum struct {
	E_1 ProjectVpcChannelInfoStatus
	E_2 ProjectVpcChannelInfoStatus
}

func GetProjectVpcChannelInfoStatusEnum() ProjectVpcChannelInfoStatusEnum {
	return ProjectVpcChannelInfoStatusEnum{
		E_1: ProjectVpcChannelInfoStatus{
			value: 1,
		}, E_2: ProjectVpcChannelInfoStatus{
			value: 2,
		},
	}
}

func (c ProjectVpcChannelInfoStatus) Value() int32 {
	return c.value
}

func (c ProjectVpcChannelInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcChannelInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ProjectVpcChannelInfoType struct {
	value string
}

type ProjectVpcChannelInfoTypeEnum struct {
	BUILTIN      ProjectVpcChannelInfoType
	MICROSERVICE ProjectVpcChannelInfoType
}

func GetProjectVpcChannelInfoTypeEnum() ProjectVpcChannelInfoTypeEnum {
	return ProjectVpcChannelInfoTypeEnum{
		BUILTIN: ProjectVpcChannelInfoType{
			value: "BUILTIN",
		},
		MICROSERVICE: ProjectVpcChannelInfoType{
			value: "MICROSERVICE",
		},
	}
}

func (c ProjectVpcChannelInfoType) Value() string {
	return c.value
}

func (c ProjectVpcChannelInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProjectVpcChannelInfoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
