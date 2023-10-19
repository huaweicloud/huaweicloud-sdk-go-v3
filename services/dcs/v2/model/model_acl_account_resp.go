package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AclAccountResp 单个ACL账号响应体
type AclAccountResp struct {

	// 账号ID
	AccountId *string `json:"account_id,omitempty"`

	// 账号名
	AccountName *string `json:"account_name,omitempty"`

	// 账号类型
	AccountType *AclAccountRespAccountType `json:"account_type,omitempty"`

	// 账号所属实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// ACL账号状态 取值范围： - CREATING：账号创建中。 - AVAILABLE：账号可用。 - CREATEFAILED：账号创建失败。 - DELETED：账号已删除。 - DELETEFAILED：账号删除失败。 - DELETING：账号删除中。 - UPDATING：账号更新中。 - ERROR：账号异常。
	Status *AclAccountRespStatus `json:"status,omitempty"`

	// 账号权限
	AccountRole *AclAccountRespAccountRole `json:"account_role,omitempty"`

	// 账号描述
	Description *string `json:"description,omitempty"`

	// 错误码（暂未使用，赋值为null）
	ErrorCode *string `json:"error_code,omitempty"`
}

func (o AclAccountResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AclAccountResp struct{}"
	}

	return strings.Join([]string{"AclAccountResp", string(data)}, " ")
}

type AclAccountRespAccountType struct {
	value string
}

type AclAccountRespAccountTypeEnum struct {
	NORMAL  AclAccountRespAccountType
	DEFAULT AclAccountRespAccountType
}

func GetAclAccountRespAccountTypeEnum() AclAccountRespAccountTypeEnum {
	return AclAccountRespAccountTypeEnum{
		NORMAL: AclAccountRespAccountType{
			value: "normal",
		},
		DEFAULT: AclAccountRespAccountType{
			value: "default",
		},
	}
}

func (c AclAccountRespAccountType) Value() string {
	return c.value
}

func (c AclAccountRespAccountType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AclAccountRespAccountType) UnmarshalJSON(b []byte) error {
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

type AclAccountRespStatus struct {
	value string
}

type AclAccountRespStatusEnum struct {
	CREATING     AclAccountRespStatus
	AVAILABLE    AclAccountRespStatus
	CREATEFAILED AclAccountRespStatus
	DELETED      AclAccountRespStatus
	DELETEFAILED AclAccountRespStatus
	DELETING     AclAccountRespStatus
	UPDATING     AclAccountRespStatus
	ERROR        AclAccountRespStatus
}

func GetAclAccountRespStatusEnum() AclAccountRespStatusEnum {
	return AclAccountRespStatusEnum{
		CREATING: AclAccountRespStatus{
			value: "CREATING",
		},
		AVAILABLE: AclAccountRespStatus{
			value: "AVAILABLE",
		},
		CREATEFAILED: AclAccountRespStatus{
			value: "CREATEFAILED",
		},
		DELETED: AclAccountRespStatus{
			value: "DELETED",
		},
		DELETEFAILED: AclAccountRespStatus{
			value: "DELETEFAILED",
		},
		DELETING: AclAccountRespStatus{
			value: "DELETING",
		},
		UPDATING: AclAccountRespStatus{
			value: "UPDATING",
		},
		ERROR: AclAccountRespStatus{
			value: "ERROR",
		},
	}
}

func (c AclAccountRespStatus) Value() string {
	return c.value
}

func (c AclAccountRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AclAccountRespStatus) UnmarshalJSON(b []byte) error {
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

type AclAccountRespAccountRole struct {
	value string
}

type AclAccountRespAccountRoleEnum struct {
	READ  AclAccountRespAccountRole
	WRITE AclAccountRespAccountRole
}

func GetAclAccountRespAccountRoleEnum() AclAccountRespAccountRoleEnum {
	return AclAccountRespAccountRoleEnum{
		READ: AclAccountRespAccountRole{
			value: "read",
		},
		WRITE: AclAccountRespAccountRole{
			value: "write",
		},
	}
}

func (c AclAccountRespAccountRole) Value() string {
	return c.value
}

func (c AclAccountRespAccountRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AclAccountRespAccountRole) UnmarshalJSON(b []byte) error {
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
