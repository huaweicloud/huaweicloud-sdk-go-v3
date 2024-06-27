package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CallIdentityPrimitiveTypeHolder struct {

	// 仅支持资源栈集权限模式为SERVICE_MANAGED时指定该参数。用于指定用户是以组织管理账号还是成员帐户中的服务委托管理员身份调用资源栈集。默认为SELF。 当资源栈集权限模式为SELF_MANAGED时，默认为SELF。 * 无论指定何种用户身份，涉及操作的资源栈集始终在组织管理账号名下。*   * `SELF` - 以组织管理账号身份调用。   * `DELEGATED_ADMIN` - 以服务委托管理员身份调用。用户的华为云账号必须在组织中已经被注册为”资源编排资源栈集服务“的委托管理员。
	CallIdentity *CallIdentityPrimitiveTypeHolderCallIdentity `json:"call_identity,omitempty"`
}

func (o CallIdentityPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CallIdentityPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"CallIdentityPrimitiveTypeHolder", string(data)}, " ")
}

type CallIdentityPrimitiveTypeHolderCallIdentity struct {
	value string
}

type CallIdentityPrimitiveTypeHolderCallIdentityEnum struct {
	SELF            CallIdentityPrimitiveTypeHolderCallIdentity
	DELEGATED_ADMIN CallIdentityPrimitiveTypeHolderCallIdentity
}

func GetCallIdentityPrimitiveTypeHolderCallIdentityEnum() CallIdentityPrimitiveTypeHolderCallIdentityEnum {
	return CallIdentityPrimitiveTypeHolderCallIdentityEnum{
		SELF: CallIdentityPrimitiveTypeHolderCallIdentity{
			value: "SELF",
		},
		DELEGATED_ADMIN: CallIdentityPrimitiveTypeHolderCallIdentity{
			value: "DELEGATED_ADMIN",
		},
	}
}

func (c CallIdentityPrimitiveTypeHolderCallIdentity) Value() string {
	return c.value
}

func (c CallIdentityPrimitiveTypeHolderCallIdentity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CallIdentityPrimitiveTypeHolderCallIdentity) UnmarshalJSON(b []byte) error {
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
