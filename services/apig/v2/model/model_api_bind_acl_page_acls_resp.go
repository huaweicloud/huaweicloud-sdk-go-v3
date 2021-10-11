package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

type ApiBindAclPageAclsResp struct {
	// ACL策略编号

	AclId *string `json:"acl_id,omitempty"`
	// ACL策略名称

	AclName *string `json:"acl_name,omitempty"`
	// ACL策略类型 - PERMIT：白名单类型 - DENY：黑名单类型

	AclType *ApiBindAclPageAclsRespAclType `json:"acl_type,omitempty"`
	// ACL策略值

	AclValue *string `json:"acl_value,omitempty"`
	// ACL策略作用的对象类型

	EntityType *ApiBindAclPageAclsRespEntityType `json:"entity_type,omitempty"`
	// 生效的环境编号

	EnvId *string `json:"env_id,omitempty"`
	// 生效的环境名称

	EnvName *string `json:"env_name,omitempty"`
	// 绑定关系编号

	BindId *string `json:"bind_id,omitempty"`
	// 绑定时间

	BindTime *sdktime.SdkTime `json:"bind_time,omitempty"`
}

func (o ApiBindAclPageAclsResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApiBindAclPageAclsResp struct{}"
	}

	return strings.Join([]string{"ApiBindAclPageAclsResp", string(data)}, " ")
}

type ApiBindAclPageAclsRespAclType struct {
	value string
}

type ApiBindAclPageAclsRespAclTypeEnum struct {
	PERMIT ApiBindAclPageAclsRespAclType
	DENY   ApiBindAclPageAclsRespAclType
}

func GetApiBindAclPageAclsRespAclTypeEnum() ApiBindAclPageAclsRespAclTypeEnum {
	return ApiBindAclPageAclsRespAclTypeEnum{
		PERMIT: ApiBindAclPageAclsRespAclType{
			value: "PERMIT",
		},
		DENY: ApiBindAclPageAclsRespAclType{
			value: "DENY",
		},
	}
}

func (c ApiBindAclPageAclsRespAclType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ApiBindAclPageAclsRespAclType) UnmarshalJSON(b []byte) error {
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

type ApiBindAclPageAclsRespEntityType struct {
	value string
}

type ApiBindAclPageAclsRespEntityTypeEnum struct {
	IP     ApiBindAclPageAclsRespEntityType
	DOMAIN ApiBindAclPageAclsRespEntityType
}

func GetApiBindAclPageAclsRespEntityTypeEnum() ApiBindAclPageAclsRespEntityTypeEnum {
	return ApiBindAclPageAclsRespEntityTypeEnum{
		IP: ApiBindAclPageAclsRespEntityType{
			value: "IP",
		},
		DOMAIN: ApiBindAclPageAclsRespEntityType{
			value: "DOMAIN",
		},
	}
}

func (c ApiBindAclPageAclsRespEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ApiBindAclPageAclsRespEntityType) UnmarshalJSON(b []byte) error {
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
