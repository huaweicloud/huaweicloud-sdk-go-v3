package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePrivateNatOption 创建私网NAT网关实例的请求体。
type CreatePrivateNatOption struct {

	// 私网NAT网关实例的名字。 私网NAT网关实例的名字仅支持数字、字母、_（下划线）、-（中划线）、中文。
	Name string `json:"name"`

	// 私网NAT网关实例的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 私网NAT网关实例的规格。 取值为： \"Small\"：小型 \"Medium\"：中型 \"Large\"：大型 \"Extra-large\"：超大型
	Spec *CreatePrivateNatOptionSpec `json:"spec,omitempty"`

	// 私网NAT网关实例所属的VPC实例。
	DownlinkVpcs []DownlinkVpcOption `json:"downlink_vpcs"`

	// 标签列表
	Tags *[]PrivateTag `json:"tags,omitempty"`

	// 企业项目ID 创建私网NAT网关实例时，关联的企业项目ID。 关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreatePrivateNatOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivateNatOption struct{}"
	}

	return strings.Join([]string{"CreatePrivateNatOption", string(data)}, " ")
}

type CreatePrivateNatOptionSpec struct {
	value string
}

type CreatePrivateNatOptionSpecEnum struct {
	SMALL       CreatePrivateNatOptionSpec
	MEDIUM      CreatePrivateNatOptionSpec
	LARGE       CreatePrivateNatOptionSpec
	EXTRA_LARGE CreatePrivateNatOptionSpec
}

func GetCreatePrivateNatOptionSpecEnum() CreatePrivateNatOptionSpecEnum {
	return CreatePrivateNatOptionSpecEnum{
		SMALL: CreatePrivateNatOptionSpec{
			value: "Small",
		},
		MEDIUM: CreatePrivateNatOptionSpec{
			value: "Medium",
		},
		LARGE: CreatePrivateNatOptionSpec{
			value: "Large",
		},
		EXTRA_LARGE: CreatePrivateNatOptionSpec{
			value: "Extra-large",
		},
	}
}

func (c CreatePrivateNatOptionSpec) Value() string {
	return c.value
}

func (c CreatePrivateNatOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivateNatOptionSpec) UnmarshalJSON(b []byte) error {
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
