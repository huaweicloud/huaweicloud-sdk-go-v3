package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// This is a auto update request Object
type UpdateUserDetailReq struct {
	// DDM实例帐号的基础权限，默认值为原DDM帐号权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT

	BaseAuthority *[]UpdateUserDetailReqBaseAuthority `json:"base_authority,omitempty"`
	// DDM实例帐号的扩展权限，默认值为空。  取值范围为：fulltableDelete、fulltableSelect、fulltableUpdate。  说明： 权限配置应该遵循如下原则：  请至少选择一个基础权限，且扩展权限对应的基础权限必须选择，对应关系如下：   - “fulltableSelect”对应“SELECT”   - “fulltableDelete”对应“DELETE”   - “fulltableUpdate”对应“UPDATE”

	ExtendAuthority *[]UpdateUserDetailReqExtendAuthority `json:"extend_authority,omitempty"`
	// DDM实例帐号的描述信息，长度不能超过256个字符。  默认值为空。

	Description *string `json:"description,omitempty"`
	// DDM实例帐号相关信息的集合。

	Databases *[]UpdateUsersDatabases `json:"databases,omitempty"`
}

func (o UpdateUserDetailReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateUserDetailReq struct{}"
	}

	return strings.Join([]string{"UpdateUserDetailReq", string(data)}, " ")
}

type UpdateUserDetailReqBaseAuthority struct {
	value string
}

type UpdateUserDetailReqBaseAuthorityEnum struct {
	CREATE UpdateUserDetailReqBaseAuthority
	DROP   UpdateUserDetailReqBaseAuthority
	ALTER  UpdateUserDetailReqBaseAuthority
	INDEX  UpdateUserDetailReqBaseAuthority
	INSERT UpdateUserDetailReqBaseAuthority
	DELETE UpdateUserDetailReqBaseAuthority
	UPDATE UpdateUserDetailReqBaseAuthority
	SELECT UpdateUserDetailReqBaseAuthority
}

func GetUpdateUserDetailReqBaseAuthorityEnum() UpdateUserDetailReqBaseAuthorityEnum {
	return UpdateUserDetailReqBaseAuthorityEnum{
		CREATE: UpdateUserDetailReqBaseAuthority{
			value: "CREATE",
		},
		DROP: UpdateUserDetailReqBaseAuthority{
			value: "DROP",
		},
		ALTER: UpdateUserDetailReqBaseAuthority{
			value: "ALTER",
		},
		INDEX: UpdateUserDetailReqBaseAuthority{
			value: "INDEX",
		},
		INSERT: UpdateUserDetailReqBaseAuthority{
			value: "INSERT",
		},
		DELETE: UpdateUserDetailReqBaseAuthority{
			value: "DELETE",
		},
		UPDATE: UpdateUserDetailReqBaseAuthority{
			value: "UPDATE",
		},
		SELECT: UpdateUserDetailReqBaseAuthority{
			value: "SELECT",
		},
	}
}

func (c UpdateUserDetailReqBaseAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateUserDetailReqBaseAuthority) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateUserDetailReqExtendAuthority struct {
	value string
}

type UpdateUserDetailReqExtendAuthorityEnum struct {
	FULLTABLE_SELECT UpdateUserDetailReqExtendAuthority
	FULLTABLE_DELETE UpdateUserDetailReqExtendAuthority
	FULLTABLE_UPDATE UpdateUserDetailReqExtendAuthority
}

func GetUpdateUserDetailReqExtendAuthorityEnum() UpdateUserDetailReqExtendAuthorityEnum {
	return UpdateUserDetailReqExtendAuthorityEnum{
		FULLTABLE_SELECT: UpdateUserDetailReqExtendAuthority{
			value: "fulltableSelect",
		},
		FULLTABLE_DELETE: UpdateUserDetailReqExtendAuthority{
			value: "fulltableDelete",
		},
		FULLTABLE_UPDATE: UpdateUserDetailReqExtendAuthority{
			value: "fulltableUpdate",
		},
	}
}

func (c UpdateUserDetailReqExtendAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateUserDetailReqExtendAuthority) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
