package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// CreateAclStrategyV2Response Response Object
type CreateAclStrategyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型。 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *CreateAclStrategyV2ResponseAclType `json:"acl_type,omitempty"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty"`

	// 对象类型： - IP - DOMAIN - DOMAIN_ID
	EntityType *string `json:"entity_type,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 更新时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateAclStrategyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAclStrategyV2Response struct{}"
	}

	return strings.Join([]string{"CreateAclStrategyV2Response", string(data)}, " ")
}

type CreateAclStrategyV2ResponseAclType struct {
	value string
}

type CreateAclStrategyV2ResponseAclTypeEnum struct {
	PERMIT CreateAclStrategyV2ResponseAclType
	DENY   CreateAclStrategyV2ResponseAclType
}

func GetCreateAclStrategyV2ResponseAclTypeEnum() CreateAclStrategyV2ResponseAclTypeEnum {
	return CreateAclStrategyV2ResponseAclTypeEnum{
		PERMIT: CreateAclStrategyV2ResponseAclType{
			value: "PERMIT",
		},
		DENY: CreateAclStrategyV2ResponseAclType{
			value: "DENY",
		},
	}
}

func (c CreateAclStrategyV2ResponseAclType) Value() string {
	return c.value
}

func (c CreateAclStrategyV2ResponseAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateAclStrategyV2ResponseAclType) UnmarshalJSON(b []byte) error {
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
