package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDetailsOfAclPolicyV2Response Response Object
type ShowDetailsOfAclPolicyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型。 - PERMIT：白名单类型 - DENY：黑名单类型
	AclType *ShowDetailsOfAclPolicyV2ResponseAclType `json:"acl_type,omitempty"`

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

func (o ShowDetailsOfAclPolicyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAclPolicyV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAclPolicyV2Response", string(data)}, " ")
}

type ShowDetailsOfAclPolicyV2ResponseAclType struct {
	value string
}

type ShowDetailsOfAclPolicyV2ResponseAclTypeEnum struct {
	PERMIT ShowDetailsOfAclPolicyV2ResponseAclType
	DENY   ShowDetailsOfAclPolicyV2ResponseAclType
}

func GetShowDetailsOfAclPolicyV2ResponseAclTypeEnum() ShowDetailsOfAclPolicyV2ResponseAclTypeEnum {
	return ShowDetailsOfAclPolicyV2ResponseAclTypeEnum{
		PERMIT: ShowDetailsOfAclPolicyV2ResponseAclType{
			value: "PERMIT",
		},
		DENY: ShowDetailsOfAclPolicyV2ResponseAclType{
			value: "DENY",
		},
	}
}

func (c ShowDetailsOfAclPolicyV2ResponseAclType) Value() string {
	return c.value
}

func (c ShowDetailsOfAclPolicyV2ResponseAclType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfAclPolicyV2ResponseAclType) UnmarshalJSON(b []byte) error {
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
