package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountAssignmentDto 创建绑定的实体类型
type AccountAssignmentDto struct {

	// 账号的唯一标识
	AccountId *string `json:"account_id,omitempty"`

	// 权限集唯一标识.
	PermissionSetId *string `json:"permission_set_id,omitempty"`

	// IAM身份中心中的一个实体身份唯一标识，例如用户或用户组
	PrincipalId *string `json:"principal_id,omitempty"`

	// 绑定的实体类型
	PrincipalType *AccountAssignmentDtoPrincipalType `json:"principal_type,omitempty"`
}

func (o AccountAssignmentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountAssignmentDto struct{}"
	}

	return strings.Join([]string{"AccountAssignmentDto", string(data)}, " ")
}

type AccountAssignmentDtoPrincipalType struct {
	value string
}

type AccountAssignmentDtoPrincipalTypeEnum struct {
	USER  AccountAssignmentDtoPrincipalType
	GROUP AccountAssignmentDtoPrincipalType
}

func GetAccountAssignmentDtoPrincipalTypeEnum() AccountAssignmentDtoPrincipalTypeEnum {
	return AccountAssignmentDtoPrincipalTypeEnum{
		USER: AccountAssignmentDtoPrincipalType{
			value: "USER",
		},
		GROUP: AccountAssignmentDtoPrincipalType{
			value: "GROUP",
		},
	}
}

func (c AccountAssignmentDtoPrincipalType) Value() string {
	return c.value
}

func (c AccountAssignmentDtoPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountAssignmentDtoPrincipalType) UnmarshalJSON(b []byte) error {
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
