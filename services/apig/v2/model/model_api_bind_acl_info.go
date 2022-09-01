package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiBindAclInfo struct {

	// ACL策略编号
	AclId *string `json:"acl_id,omitempty" xml:"acl_id"`

	// ACL策略名称
	AclName *string `json:"acl_name,omitempty" xml:"acl_name"`

	// ACL策略作用的对象类型
	EntityType *ApiBindAclInfoEntityType `json:"entity_type,omitempty" xml:"entity_type"`

	// ACL策略类型 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *ApiBindAclInfoAclType `json:"acl_type,omitempty" xml:"acl_type"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty" xml:"acl_value"`

	// 生效的环境编号
	EnvId *string `json:"env_id,omitempty" xml:"env_id"`

	// 生效的环境名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// 绑定关系编号
	BindId *string `json:"bind_id,omitempty" xml:"bind_id"`

	// 绑定时间
	BindTime *sdktime.SdkTime `json:"bind_time,omitempty" xml:"bind_time"`
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
