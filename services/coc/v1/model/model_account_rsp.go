package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccountRsp 账号返回体
type AccountRsp struct {

	// id
	Id string `json:"id"`

	// 名称
	Name string `json:"name"`

	// 账号类型 NORMAL 普通账号，停用状态 DELEGATOR 委托人 DELEGATEE 被委托人
	Type *AccountRspType `json:"type,omitempty"`

	// 被委托人id
	Delegatee *string `json:"delegatee,omitempty"`

	// 委托id
	AgencyId *string `json:"agency_id,omitempty"`

	// 托管账号Id
	Delegator *string `json:"delegator,omitempty"`

	// 创建人
	CreateUserName string `json:"create_user_name"`

	// PENDING_AUTHORIZATION, 待授权,ENABLED 启用, DISABLED 停用
	Status AccountRspStatus `json:"status"`

	// 创建userId
	CreateUserId string `json:"create_user_id"`

	// 创建时间
	CreateTime string `json:"create_time"`

	// 编辑人
	UpdateUserName *string `json:"update_user_name,omitempty"`

	// 编辑人user_id
	UpdateUserId *string `json:"update_user_id,omitempty"`

	// 编辑时间，期初和创建时间一样
	UpdateTime string `json:"update_time"`
}

func (o AccountRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountRsp struct{}"
	}

	return strings.Join([]string{"AccountRsp", string(data)}, " ")
}

type AccountRspType struct {
	value string
}

type AccountRspTypeEnum struct {
	NORMAL    AccountRspType
	DELEGATOR AccountRspType
	DELEGATEE AccountRspType
}

func GetAccountRspTypeEnum() AccountRspTypeEnum {
	return AccountRspTypeEnum{
		NORMAL: AccountRspType{
			value: "NORMAL",
		},
		DELEGATOR: AccountRspType{
			value: "DELEGATOR",
		},
		DELEGATEE: AccountRspType{
			value: "DELEGATEE",
		},
	}
}

func (c AccountRspType) Value() string {
	return c.value
}

func (c AccountRspType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountRspType) UnmarshalJSON(b []byte) error {
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

type AccountRspStatus struct {
	value string
}

type AccountRspStatusEnum struct {
	PENDING_AUTHORIZATION AccountRspStatus
	ENABLED               AccountRspStatus
	DISABLED              AccountRspStatus
}

func GetAccountRspStatusEnum() AccountRspStatusEnum {
	return AccountRspStatusEnum{
		PENDING_AUTHORIZATION: AccountRspStatus{
			value: "PENDING_AUTHORIZATION",
		},
		ENABLED: AccountRspStatus{
			value: "ENABLED",
		},
		DISABLED: AccountRspStatus{
			value: "DISABLED",
		},
	}
}

func (c AccountRspStatus) Value() string {
	return c.value
}

func (c AccountRspStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccountRspStatus) UnmarshalJSON(b []byte) error {
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
