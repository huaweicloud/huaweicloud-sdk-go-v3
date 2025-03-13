package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PermissionApprovalDetailDtoProposers struct {

	// 申请人id
	Id *string `json:"id,omitempty"`

	// 申请人名称
	Name *string `json:"name,omitempty"`

	// 申请人类型,USER,USER_GROUP,ROLE,WORKSPACE_ACCOUNT
	Type *PermissionApprovalDetailDtoProposersType `json:"type,omitempty"`
}

func (o PermissionApprovalDetailDtoProposers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionApprovalDetailDtoProposers struct{}"
	}

	return strings.Join([]string{"PermissionApprovalDetailDtoProposers", string(data)}, " ")
}

type PermissionApprovalDetailDtoProposersType struct {
	value string
}

type PermissionApprovalDetailDtoProposersTypeEnum struct {
	USER              PermissionApprovalDetailDtoProposersType
	USER_GROUP        PermissionApprovalDetailDtoProposersType
	ROLE              PermissionApprovalDetailDtoProposersType
	WORKSPACE_ACCOUNT PermissionApprovalDetailDtoProposersType
}

func GetPermissionApprovalDetailDtoProposersTypeEnum() PermissionApprovalDetailDtoProposersTypeEnum {
	return PermissionApprovalDetailDtoProposersTypeEnum{
		USER: PermissionApprovalDetailDtoProposersType{
			value: "USER",
		},
		USER_GROUP: PermissionApprovalDetailDtoProposersType{
			value: "USER_GROUP",
		},
		ROLE: PermissionApprovalDetailDtoProposersType{
			value: "ROLE",
		},
		WORKSPACE_ACCOUNT: PermissionApprovalDetailDtoProposersType{
			value: "WORKSPACE_ACCOUNT",
		},
	}
}

func (c PermissionApprovalDetailDtoProposersType) Value() string {
	return c.value
}

func (c PermissionApprovalDetailDtoProposersType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PermissionApprovalDetailDtoProposersType) UnmarshalJSON(b []byte) error {
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
