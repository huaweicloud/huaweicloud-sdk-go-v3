package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListProjectRolePermissionsResponse Response Object
type ListProjectRolePermissionsResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListProjectRolePermissionsResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**: 权限列表。 **取值范围**: 不涉及。
	Result         *[]ProjectRolePermissionDo `json:"result,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListProjectRolePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectRolePermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectRolePermissionsResponse", string(data)}, " ")
}

type ListProjectRolePermissionsResponseStatus struct {
	value string
}

type ListProjectRolePermissionsResponseStatusEnum struct {
	SUCCESS ListProjectRolePermissionsResponseStatus
	ERROR   ListProjectRolePermissionsResponseStatus
}

func GetListProjectRolePermissionsResponseStatusEnum() ListProjectRolePermissionsResponseStatusEnum {
	return ListProjectRolePermissionsResponseStatusEnum{
		SUCCESS: ListProjectRolePermissionsResponseStatus{
			value: "success",
		},
		ERROR: ListProjectRolePermissionsResponseStatus{
			value: "error",
		},
	}
}

func (c ListProjectRolePermissionsResponseStatus) Value() string {
	return c.value
}

func (c ListProjectRolePermissionsResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListProjectRolePermissionsResponseStatus) UnmarshalJSON(b []byte) error {
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
