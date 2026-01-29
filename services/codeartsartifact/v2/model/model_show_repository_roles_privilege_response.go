package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryRolesPrivilegeResponse Response Object
type ShowRepositoryRolesPrivilegeResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowRepositoryRolesPrivilegeResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 查询结果。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 无。
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRepositoryRolesPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryRolesPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryRolesPrivilegeResponse", string(data)}, " ")
}

type ShowRepositoryRolesPrivilegeResponseStatus struct {
	value string
}

type ShowRepositoryRolesPrivilegeResponseStatusEnum struct {
	SUCCESS ShowRepositoryRolesPrivilegeResponseStatus
	ERROR   ShowRepositoryRolesPrivilegeResponseStatus
}

func GetShowRepositoryRolesPrivilegeResponseStatusEnum() ShowRepositoryRolesPrivilegeResponseStatusEnum {
	return ShowRepositoryRolesPrivilegeResponseStatusEnum{
		SUCCESS: ShowRepositoryRolesPrivilegeResponseStatus{
			value: "success",
		},
		ERROR: ShowRepositoryRolesPrivilegeResponseStatus{
			value: "error",
		},
	}
}

func (c ShowRepositoryRolesPrivilegeResponseStatus) Value() string {
	return c.value
}

func (c ShowRepositoryRolesPrivilegeResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryRolesPrivilegeResponseStatus) UnmarshalJSON(b []byte) error {
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
