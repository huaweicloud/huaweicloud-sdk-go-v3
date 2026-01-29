package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateRepoRolesPrivilegeResponse Response Object
type UpdateRepoRolesPrivilegeResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *UpdateRepoRolesPrivilegeResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// 参数解释: 权限信息。 取值范围: 不涉及。
	Result         *[]Privilege `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateRepoRolesPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepoRolesPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepoRolesPrivilegeResponse", string(data)}, " ")
}

type UpdateRepoRolesPrivilegeResponseStatus struct {
	value string
}

type UpdateRepoRolesPrivilegeResponseStatusEnum struct {
	SUCCESS UpdateRepoRolesPrivilegeResponseStatus
	ERROR   UpdateRepoRolesPrivilegeResponseStatus
}

func GetUpdateRepoRolesPrivilegeResponseStatusEnum() UpdateRepoRolesPrivilegeResponseStatusEnum {
	return UpdateRepoRolesPrivilegeResponseStatusEnum{
		SUCCESS: UpdateRepoRolesPrivilegeResponseStatus{
			value: "success",
		},
		ERROR: UpdateRepoRolesPrivilegeResponseStatus{
			value: "error",
		},
	}
}

func (c UpdateRepoRolesPrivilegeResponseStatus) Value() string {
	return c.value
}

func (c UpdateRepoRolesPrivilegeResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRepoRolesPrivilegeResponseStatus) UnmarshalJSON(b []byte) error {
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
