/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type UpdatePublicipOption struct {
	// 功能说明：公网IP的名称
	Alias *string `json:"alias,omitempty"`
	// 功能说明：公网IP的描述信息 取值范围：0-256长度的字符串，不支持特殊字符<>
	Description *string `json:"description,omitempty"`
	// 功能说明：端口所属实例类型 取值范围：PORT、NATGW、VPN、ELB、null 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例； associate_instance_type和associate_instance_id都为null时解绑实例 约束：双栈公网IP不允许修改绑定的实例
	AssociateInstanceType *UpdatePublicipOptionAssociateInstanceType `json:"associate_instance_type,omitempty"`
	// 功能说明：端口所属实例ID，例如RDS实例ID 约束：associate_instance_type和associate_instance_id都不为空时表示绑定实例； associate_instance_type和associate_instance_id都为null时解绑实例 约束：双栈公网IP不允许修改绑定的实例
	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`
}

func (o UpdatePublicipOption) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePublicipOption", string(data)}, " ")
}

type UpdatePublicipOptionAssociateInstanceType struct {
	value string
}

type UpdatePublicipOptionAssociateInstanceTypeEnum struct {
	PORT  UpdatePublicipOptionAssociateInstanceType
	NATGW UpdatePublicipOptionAssociateInstanceType
	ELB   UpdatePublicipOptionAssociateInstanceType
	EMPTY UpdatePublicipOptionAssociateInstanceType
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
		EMPTY: UpdatePublicipOptionAssociateInstanceType{
			value: "",
		},
	}
}

func (c UpdatePublicipOptionAssociateInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdatePublicipOptionAssociateInstanceType) UnmarshalJSON(b []byte) error {
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
