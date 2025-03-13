package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySecurityAdminResponse Response Object
type ModifySecurityAdminResponse struct {

	// 管理员类型, USER:用户, USER_GROUP:用户组
	ManagerType *ModifySecurityAdminResponseManagerType `json:"manager_type,omitempty"`

	// 管理员id, 管理员类型为用户时, 请传入iam用户id; 管理员类型为用户组时, 请传入iam用户组id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称, 管理员类型为用户时, 请传入iam用户名称; 管理员类型为用户组时, 请传入iam用户组名称
	ManagerName    *string `json:"manager_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifySecurityAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityAdminResponse struct{}"
	}

	return strings.Join([]string{"ModifySecurityAdminResponse", string(data)}, " ")
}

type ModifySecurityAdminResponseManagerType struct {
	value string
}

type ModifySecurityAdminResponseManagerTypeEnum struct {
	USER       ModifySecurityAdminResponseManagerType
	USER_GROUP ModifySecurityAdminResponseManagerType
}

func GetModifySecurityAdminResponseManagerTypeEnum() ModifySecurityAdminResponseManagerTypeEnum {
	return ModifySecurityAdminResponseManagerTypeEnum{
		USER: ModifySecurityAdminResponseManagerType{
			value: "USER",
		},
		USER_GROUP: ModifySecurityAdminResponseManagerType{
			value: "USER_GROUP",
		},
	}
}

func (c ModifySecurityAdminResponseManagerType) Value() string {
	return c.value
}

func (c ModifySecurityAdminResponseManagerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySecurityAdminResponseManagerType) UnmarshalJSON(b []byte) error {
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
