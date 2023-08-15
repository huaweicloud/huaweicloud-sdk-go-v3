package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiBindAclInfo struct {

	// ACL策略编号
	AclId *string `json:"acl_id,omitempty"`

	// ACL策略名称
	AclName *string `json:"acl_name,omitempty"`

	// ACL策略作用的对象类型
	EntityType *ApiBindAclInfoEntityType `json:"entity_type,omitempty"`

	// ACL策略类型 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *ApiBindAclInfoAclType `json:"acl_type,omitempty"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty"`

	// 生效的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// 生效的环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 绑定关系编号
	BindId *string `json:"bind_id,omitempty"`

	// 绑定时间
	BindTime *sdktime.SdkTime `json:"bind_time,omitempty"`
}

func (o ApiBindAclInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiBindAclInfo struct{}"
	}

	return strings.Join([]string{"ApiBindAclInfo", string(data)}, " ")
}

type ApiBindAclInfoEntityType struct {
	value string
}

type ApiBindAclInfoEntityTypeEnum struct {
	IP     ApiBindAclInfoEntityType
	DOMAIN ApiBindAclInfoEntityType
}

func GetApiBindAclInfoEntityTypeEnum() ApiBindAclInfoEntityTypeEnum {
	return ApiBindAclInfoEntityTypeEnum{
		IP: ApiBindAclInfoEntityType{
			value: "IP",
		},
		DOMAIN: ApiBindAclInfoEntityType{
			value: "DOMAIN",
		},
	}
}

func (c ApiBindAclInfoEntityType) Value() string {
	return c.value
}

func (c ApiBindAclInfoEntityType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBindAclInfoEntityType) UnmarshalJSON(b []byte) error {
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

type ApiBindAclInfoAclType struct {
	value string
}

type ApiBindAclInfoAclTypeEnum struct {
	PERMIT ApiBindAclInfoAclType
	DENY   ApiBindAclInfoAclType
}

func GetApiBindAclInfoAclTypeEnum() ApiBindAclInfoAclTypeEnum {
	return ApiBindAclInfoAclTypeEnum{
		PERMIT: ApiBindAclInfoAclType{
			value: "PERMIT",
		},
		DENY: ApiBindAclInfoAclType{
			value: "DENY",
		},
	}
}

func (c ApiBindAclInfoAclType) Value() string {
	return c.value
}

func (c ApiBindAclInfoAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiBindAclInfoAclType) UnmarshalJSON(b []byte) error {
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
