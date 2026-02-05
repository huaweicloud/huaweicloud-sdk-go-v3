package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ApiAclInfoWithBindNum struct {

	// ACL策略名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型。 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *ApiAclInfoWithBindNumAclType `json:"acl_type,omitempty"`

	// ACL策略的值
	AclValue *string `json:"acl_value,omitempty"`

	// 绑定的API数量
	BindNum *int32 `json:"bind_num,omitempty"`

	// 对象类型 - IP - DOMAIN - DOMAIN_ID
	EntityType *string `json:"entity_type,omitempty"`

	// ACL策略编号
	Id *string `json:"id,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o ApiAclInfoWithBindNum) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAclInfoWithBindNum struct{}"
	}

	return strings.Join([]string{"ApiAclInfoWithBindNum", string(data)}, " ")
}

type ApiAclInfoWithBindNumAclType struct {
	value string
}

type ApiAclInfoWithBindNumAclTypeEnum struct {
	PERMIT ApiAclInfoWithBindNumAclType
	DENY   ApiAclInfoWithBindNumAclType
}

func GetApiAclInfoWithBindNumAclTypeEnum() ApiAclInfoWithBindNumAclTypeEnum {
	return ApiAclInfoWithBindNumAclTypeEnum{
		PERMIT: ApiAclInfoWithBindNumAclType{
			value: "PERMIT",
		},
		DENY: ApiAclInfoWithBindNumAclType{
			value: "DENY",
		},
	}
}

func (c ApiAclInfoWithBindNumAclType) Value() string {
	return c.value
}

func (c ApiAclInfoWithBindNumAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiAclInfoWithBindNumAclType) UnmarshalJSON(b []byte) error {
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
