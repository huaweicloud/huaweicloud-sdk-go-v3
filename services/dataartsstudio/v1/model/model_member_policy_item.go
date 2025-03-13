package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MemberPolicyItem struct {

	// 成员id
	MemberId string `json:"member_id"`

	// 成员名称
	MemberName string `json:"member_name"`

	// 成员类型:USER,USER_GROUP,WORKSPACE_ROLE，分别代表空间用户、空间用户组、空间角色
	MemberType *MemberPolicyItemMemberType `json:"member_type,omitempty"`
}

func (o MemberPolicyItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberPolicyItem struct{}"
	}

	return strings.Join([]string{"MemberPolicyItem", string(data)}, " ")
}

type MemberPolicyItemMemberType struct {
	value string
}

type MemberPolicyItemMemberTypeEnum struct {
	USER           MemberPolicyItemMemberType
	USER_GROUP     MemberPolicyItemMemberType
	WORKSPACE_ROLE MemberPolicyItemMemberType
}

func GetMemberPolicyItemMemberTypeEnum() MemberPolicyItemMemberTypeEnum {
	return MemberPolicyItemMemberTypeEnum{
		USER: MemberPolicyItemMemberType{
			value: "USER",
		},
		USER_GROUP: MemberPolicyItemMemberType{
			value: "USER_GROUP",
		},
		WORKSPACE_ROLE: MemberPolicyItemMemberType{
			value: "WORKSPACE_ROLE",
		},
	}
}

func (c MemberPolicyItemMemberType) Value() string {
	return c.value
}

func (c MemberPolicyItemMemberType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MemberPolicyItemMemberType) UnmarshalJSON(b []byte) error {
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
