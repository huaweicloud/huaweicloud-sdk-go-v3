package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type FindingPrincipal struct {

	// 主体身份的标识符。
	Identifier string `json:"identifier"`

	// 主体身份的类型。 - all_principal：所有主体 - account：账号 - all_user_in_account： 账号下所有用户 - all_agency_in_account： 账号下所有委托 - all_identity_provider_in_account：账号下所有身份提供商 - specific_user：特定用户 - specific_agency：特定委托 - specific_group：特定用户组 - specific_identity_provider：特定身份提供商
	Type FindingPrincipalType `json:"type"`
}

func (o FindingPrincipal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FindingPrincipal struct{}"
	}

	return strings.Join([]string{"FindingPrincipal", string(data)}, " ")
}

type FindingPrincipalType struct {
	value string
}

type FindingPrincipalTypeEnum struct {
	ALL_PRINCIPAL                    FindingPrincipalType
	ACCOUNT                          FindingPrincipalType
	ALL_USER_IN_ACCOUNT              FindingPrincipalType
	ALL_AGENCY_IN_ACCOUNT            FindingPrincipalType
	ALL_IDENTITY_PROVIDER_IN_ACCOUNT FindingPrincipalType
	SPECIFIC_USER                    FindingPrincipalType
	SPECIFIC_AGENCY                  FindingPrincipalType
	SPECIFIC_GROUP                   FindingPrincipalType
	SPECIFIC_IDENTITY_PROVIDER       FindingPrincipalType
}

func GetFindingPrincipalTypeEnum() FindingPrincipalTypeEnum {
	return FindingPrincipalTypeEnum{
		ALL_PRINCIPAL: FindingPrincipalType{
			value: "all_principal",
		},
		ACCOUNT: FindingPrincipalType{
			value: "account",
		},
		ALL_USER_IN_ACCOUNT: FindingPrincipalType{
			value: "all_user_in_account",
		},
		ALL_AGENCY_IN_ACCOUNT: FindingPrincipalType{
			value: "all_agency_in_account",
		},
		ALL_IDENTITY_PROVIDER_IN_ACCOUNT: FindingPrincipalType{
			value: "all_identity_provider_in_account",
		},
		SPECIFIC_USER: FindingPrincipalType{
			value: "specific_user",
		},
		SPECIFIC_AGENCY: FindingPrincipalType{
			value: "specific_agency",
		},
		SPECIFIC_GROUP: FindingPrincipalType{
			value: "specific_group",
		},
		SPECIFIC_IDENTITY_PROVIDER: FindingPrincipalType{
			value: "specific_identity_provider",
		},
	}
}

func (c FindingPrincipalType) Value() string {
	return c.value
}

func (c FindingPrincipalType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingPrincipalType) UnmarshalJSON(b []byte) error {
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
