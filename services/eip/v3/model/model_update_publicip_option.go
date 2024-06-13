package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdatePublicipOption struct {

	// 功能说明：公网IP的名称。
	Alias *string `json:"alias,omitempty"`

	// 功能说明：公网IP的描述信息 取值范围：0-256长度的字符串，不支持特殊字符<>
	Description *string `json:"description,omitempty"`

	// 功能说明：端口所属实例类型 取值范围：PORT、NATGW、VPN、ELB、null 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例。 约束：associate_instance_type和associate_instance_id都为null时表示解绑实例，通过APIE调用需要切换为文本输入方式输入null值，可参考解绑请求实例。 约束：双栈公网IP不允许修改绑定的实例。
	AssociateInstanceType *UpdatePublicipOptionAssociateInstanceType `json:"associate_instance_type,omitempty"`

	// 功能说明：端口所属实例ID，例如RDS实例ID 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例。 约束：associate_instance_type和associate_instance_id都为null时表示解绑实例，通过APIE调用需要切换为文本输入方式输入null值，可参考解绑请求实例。 约束：双栈公网IP不允许修改绑定的实例。
	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`
}

func (o UpdatePublicipOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipOption struct{}"
	}

	return strings.Join([]string{"UpdatePublicipOption", string(data)}, " ")
}

type UpdatePublicipOptionAssociateInstanceType struct {
	value string
}

type UpdatePublicipOptionAssociateInstanceTypeEnum struct {
	PORT  UpdatePublicipOptionAssociateInstanceType
	NATGW UpdatePublicipOptionAssociateInstanceType
	ELB   UpdatePublicipOptionAssociateInstanceType
	VPN   UpdatePublicipOptionAssociateInstanceType
}

func GetUpdatePublicipOptionAssociateInstanceTypeEnum() UpdatePublicipOptionAssociateInstanceTypeEnum {
	return UpdatePublicipOptionAssociateInstanceTypeEnum{
		PORT: UpdatePublicipOptionAssociateInstanceType{
			value: "PORT",
		},
		NATGW: UpdatePublicipOptionAssociateInstanceType{
			value: "NATGW",
		},
		ELB: UpdatePublicipOptionAssociateInstanceType{
			value: "ELB",
		},
		VPN: UpdatePublicipOptionAssociateInstanceType{
			value: "VPN",
		},
	}
}

func (c UpdatePublicipOptionAssociateInstanceType) Value() string {
	return c.value
}

func (c UpdatePublicipOptionAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePublicipOptionAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
