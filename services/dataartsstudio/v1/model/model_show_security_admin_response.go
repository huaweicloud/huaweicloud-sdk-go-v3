package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSecurityAdminResponse Response Object
type ShowSecurityAdminResponse struct {

	// 管理员类型, USER:用户, USER_GROUP:用户组
	ManagerType *ShowSecurityAdminResponseManagerType `json:"manager_type,omitempty"`

	// 管理员id, 管理员类型为用户时, 请传入iam用户id; 管理员类型为用户组时, 请传入iam用户组id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称, 管理员类型为用户时, 请传入iam用户名称; 管理员类型为用户组时, 请传入iam用户组名称
	ManagerName    *string `json:"manager_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSecurityAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityAdminResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityAdminResponse", string(data)}, " ")
}

type ShowSecurityAdminResponseManagerType struct {
	value string
}

type ShowSecurityAdminResponseManagerTypeEnum struct {
	USER       ShowSecurityAdminResponseManagerType
	USER_GROUP ShowSecurityAdminResponseManagerType
}

func GetShowSecurityAdminResponseManagerTypeEnum() ShowSecurityAdminResponseManagerTypeEnum {
	return ShowSecurityAdminResponseManagerTypeEnum{
		USER: ShowSecurityAdminResponseManagerType{
			value: "USER",
		},
		USER_GROUP: ShowSecurityAdminResponseManagerType{
			value: "USER_GROUP",
		},
	}
}

func (c ShowSecurityAdminResponseManagerType) Value() string {
	return c.value
}

func (c ShowSecurityAdminResponseManagerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSecurityAdminResponseManagerType) UnmarshalJSON(b []byte) error {
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
