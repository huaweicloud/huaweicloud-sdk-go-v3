package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateAclStrategyV2Response Response Object
type UpdateAclStrategyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型。 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *UpdateAclStrategyV2ResponseAclType `json:"acl_type,omitempty"`

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

func (o UpdateAclStrategyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclStrategyV2Response struct{}"
	}

	return strings.Join([]string{"UpdateAclStrategyV2Response", string(data)}, " ")
}

type UpdateAclStrategyV2ResponseAclType struct {
	value string
}

type UpdateAclStrategyV2ResponseAclTypeEnum struct {
	PERMIT UpdateAclStrategyV2ResponseAclType
	DENY   UpdateAclStrategyV2ResponseAclType
}

func GetUpdateAclStrategyV2ResponseAclTypeEnum() UpdateAclStrategyV2ResponseAclTypeEnum {
	return UpdateAclStrategyV2ResponseAclTypeEnum{
		PERMIT: UpdateAclStrategyV2ResponseAclType{
			value: "PERMIT",
		},
		DENY: UpdateAclStrategyV2ResponseAclType{
			value: "DENY",
		},
	}
}

func (c UpdateAclStrategyV2ResponseAclType) Value() string {
	return c.value
}

func (c UpdateAclStrategyV2ResponseAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateAclStrategyV2ResponseAclType) UnmarshalJSON(b []byte) error {
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
