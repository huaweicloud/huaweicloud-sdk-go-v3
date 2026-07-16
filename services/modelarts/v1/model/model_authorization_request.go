package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AuthorizationRequest struct {

	// **参数解释**：用户ID，获取方法请参见[获取用户ID和名称](modelarts_03_0006.xml)。当user_id为all时，表示对所有IAM子用户进行授权，如果已有部分用户已授权，则更新授权。仅当授权类型为委托时，需要该字段。 **约束限制**：不涉及。 **取值范围**：字符串长度在3到32个字符之间，支持大小写字母、数字、中划线。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：授权类型。推荐使用委托方式。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - agency：委托 - credential：访问密钥（AK/SK）  **默认取值**：不涉及。
	Type *AuthorizationRequestType `json:"type,omitempty"`

	// **参数解释**：授权内容。 **约束限制**： - 当授权类型是委托，该字段为委托名称。 - 当授权类型是访问密钥，该字段为访问密钥ID（AK）。  **取值范围**：长度限制64个字符。 **默认取值**：不涉及。
	Content string `json:"content"`

	// **参数解释**：秘密访问密钥（SK）。 **约束限制**：仅当授权类型为访问密钥时，需要该字段。 **取值范围**：字符串长度为40，支持大小写字母、数字。 **默认取值**：不涉及。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释**：用户名。 **约束限制**：当user_id为all-users时，显示为所有用户。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：用户类型。 **约束限制**：不涉及。 **取值范围**：枚举类型，取值如下： - iam：授权对象类型是IAM子用户，必传字段user_id。 - federate：授权对象类型是联邦用户，必传字段user_name，user_id不传。 - federation-group：授权对象类型是联邦用户组，必传字段user_id，值为联邦用户组的id。 - grant：授权对象类型是委托用户，必传字段user_id，值为委托用户的委托id。 - all-users：授权对象类型是所有用户，必传字段user_id值是all。  **默认取值**：IAM。
	UserType *string `json:"user_type,omitempty"`
}

func (o AuthorizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizationRequest struct{}"
	}

	return strings.Join([]string{"AuthorizationRequest", string(data)}, " ")
}

type AuthorizationRequestType struct {
	value string
}

type AuthorizationRequestTypeEnum struct {
	AGENCY     AuthorizationRequestType
	CREDENTIAL AuthorizationRequestType
}

func GetAuthorizationRequestTypeEnum() AuthorizationRequestTypeEnum {
	return AuthorizationRequestTypeEnum{
		AGENCY: AuthorizationRequestType{
			value: "agency",
		},
		CREDENTIAL: AuthorizationRequestType{
			value: "credential",
		},
	}
}

func (c AuthorizationRequestType) Value() string {
	return c.value
}

func (c AuthorizationRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizationRequestType) UnmarshalJSON(b []byte) error {
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
