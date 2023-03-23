package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 私网NAT网关实例的响应体。
type PrivateNat struct {

	// 私网NAT网关实例的ID。
	Id string `json:"id"`

	// 项目的ID。
	ProjectId string `json:"project_id"`

	// 私网NAT网关实例的名字。
	Name string `json:"name"`

	// 私网NAT网关实例的描述。
	Description string `json:"description"`

	// 私网NAT网关实例的规格。 取值为： \"Small\"：小型 \"Medium\"：中型 \"Large\"：大型 \"Extra-large\"：超大型
	Spec PrivateNatSpec `json:"spec"`

	// 私网NAT网关实例的状态。 取值为： \"ACTIVE\"：正常运行 \"FROZEN\"：冻结
	Status PrivateNatStatus `json:"status"`

	// 私网NAT网关实例的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 私网NAT网关实例的更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 私网NAT网关实例所属的VPC实例。
	DownlinkVpcs []DownlinkVpc `json:"downlink_vpcs"`

	// 标签列表。
	Tags *[]PrivateTag `json:"tags,omitempty"`

	// 企业项目ID。 创建私网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o PrivateNat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateNat struct{}"
	}

	return strings.Join([]string{"PrivateNat", string(data)}, " ")
}

type PrivateNatSpec struct {
	value string
}

type PrivateNatSpecEnum struct {
	SMALL       PrivateNatSpec
	MEDIUM      PrivateNatSpec
	LARGE       PrivateNatSpec
	EXTRA_LARGE PrivateNatSpec
}

func GetPrivateNatSpecEnum() PrivateNatSpecEnum {
	return PrivateNatSpecEnum{
		SMALL: PrivateNatSpec{
			value: "Small",
		},
		MEDIUM: PrivateNatSpec{
			value: "Medium",
		},
		LARGE: PrivateNatSpec{
			value: "Large",
		},
		EXTRA_LARGE: PrivateNatSpec{
			value: "Extra-large",
		},
	}
}

func (c PrivateNatSpec) Value() string {
	return c.value
}

func (c PrivateNatSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateNatSpec) UnmarshalJSON(b []byte) error {
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

type PrivateNatStatus struct {
	value string
}

type PrivateNatStatusEnum struct {
	ACTIVE PrivateNatStatus
	FROZEN PrivateNatStatus
}

func GetPrivateNatStatusEnum() PrivateNatStatusEnum {
	return PrivateNatStatusEnum{
		ACTIVE: PrivateNatStatus{
			value: "ACTIVE",
		},
		FROZEN: PrivateNatStatus{
			value: "FROZEN",
		},
	}
}

func (c PrivateNatStatus) Value() string {
	return c.value
}

func (c PrivateNatStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateNatStatus) UnmarshalJSON(b []byte) error {
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
