package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// This is a auto update request Object
type UpdateUserDetailReq struct {
	// DDM实例帐号的基础权限，默认值为原DDM帐号权限。  取值为：CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT

	BaseAuthority *[]UpdateUserDetailReqBaseAuthority `json:"base_authority,omitempty"`
	// DDM实例帐号的描述信息，长度不能超过256个字符。  默认值为空。

	Description *string `json:"description,omitempty"`
	// DDM实例帐号相关信息的集合。

	Databases *[]UpdateUsersDatabases `json:"databases,omitempty"`
}

func (o UpdateUserDetailReq) String() string {
	data, err := utils.Marshal(o)
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
	return utils.Marshal(c.value)
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
