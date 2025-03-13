package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DlsAdmin struct {

	// 管理员类型, USER:用户, USER_GROUP:用户组
	ManagerType *DlsAdminManagerType `json:"manager_type,omitempty"`

	// 管理员id, 管理员类型为用户时, 请传入iam用户id; 管理员类型为用户组时, 请传入iam用户组id
	ManagerId *string `json:"manager_id,omitempty"`

	// 管理员名称, 管理员类型为用户时, 请传入iam用户名称; 管理员类型为用户组时, 请传入iam用户组名称
	ManagerName *string `json:"manager_name,omitempty"`
}

func (o DlsAdmin) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DlsAdmin struct{}"
	}

	return strings.Join([]string{"DlsAdmin", string(data)}, " ")
}

type DlsAdminManagerType struct {
	value string
}

type DlsAdminManagerTypeEnum struct {
	USER       DlsAdminManagerType
	USER_GROUP DlsAdminManagerType
}

func GetDlsAdminManagerTypeEnum() DlsAdminManagerTypeEnum {
	return DlsAdminManagerTypeEnum{
		USER: DlsAdminManagerType{
			value: "USER",
		},
		USER_GROUP: DlsAdminManagerType{
			value: "USER_GROUP",
		},
	}
}

func (c DlsAdminManagerType) Value() string {
	return c.value
}

func (c DlsAdminManagerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DlsAdminManagerType) UnmarshalJSON(b []byte) error {
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
