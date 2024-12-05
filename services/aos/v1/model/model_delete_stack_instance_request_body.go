package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DeleteStackInstanceRequestBody struct {

	// 资源栈集（stack_set）的唯一ID。  此ID由资源编排服务在生成资源栈集的时候生成，为UUID。  由于资源栈集名仅仅在同一时间下唯一，即用户允许先生成一个叫HelloWorld的资源栈集，删除，再重新创建一个同名资源栈集。  对于团队并行开发，用户可能希望确保，当前我操作的资源栈集就是我以为的那个，而不是被其他队友删除后创建的同名资源栈集。因此，使用ID就可以做到强匹配。  资源编排服务保证每次创建的资源栈集所对应的ID都不相同，更新不会影响ID。如果给予的stack_set_id和当前资源栈集的ID不一致，则返回400
	StackSetId *string `json:"stack_set_id,omitempty"`

	DeploymentTargets *DeploymentTargets `json:"deployment_targets"`

	OperationPreferences *OperationPreferences `json:"operation_preferences,omitempty"`

	// 仅支持资源栈集权限模式为SERVICE_MANAGED时指定该参数。用于指定用户是以组织管理账号还是成员账号中的服务委托管理员身份调用资源栈集。默认为SELF。 当资源栈集权限模式为SELF_MANAGED时，默认为SELF。 * 无论指定何种用户身份，涉及操作的资源栈集始终在组织管理账号名下。*   * `SELF` - 以组织管理账号身份调用。   * `DELEGATED_ADMIN` - 以服务委托管理员身份调用。用户的华为云账号必须在组织中已经被注册为”资源编排资源栈集服务“的委托管理员。
	CallIdentity *DeleteStackInstanceRequestBodyCallIdentity `json:"call_identity,omitempty"`
}

func (o DeleteStackInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteStackInstanceRequestBody", string(data)}, " ")
}

type DeleteStackInstanceRequestBodyCallIdentity struct {
	value string
}

type DeleteStackInstanceRequestBodyCallIdentityEnum struct {
	SELF            DeleteStackInstanceRequestBodyCallIdentity
	DELEGATED_ADMIN DeleteStackInstanceRequestBodyCallIdentity
}

func GetDeleteStackInstanceRequestBodyCallIdentityEnum() DeleteStackInstanceRequestBodyCallIdentityEnum {
	return DeleteStackInstanceRequestBodyCallIdentityEnum{
		SELF: DeleteStackInstanceRequestBodyCallIdentity{
			value: "SELF",
		},
		DELEGATED_ADMIN: DeleteStackInstanceRequestBodyCallIdentity{
			value: "DELEGATED_ADMIN",
		},
	}
}

func (c DeleteStackInstanceRequestBodyCallIdentity) Value() string {
	return c.value
}

func (c DeleteStackInstanceRequestBodyCallIdentity) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteStackInstanceRequestBodyCallIdentity) UnmarshalJSON(b []byte) error {
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
