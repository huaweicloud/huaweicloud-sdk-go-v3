package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AbuseScope 算力滥用统计范围
type AbuseScope struct {

	// 统计范围类型。all_users: 所有用户, single_user: 单个用户
	Type *AbuseScopeType `json:"type,omitempty"`

	// 用户标识字段名，当前支持extra.end_user
	UserIdentifier *string `json:"user_identifier,omitempty"`
}

func (o AbuseScope) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AbuseScope struct{}"
	}

	return strings.Join([]string{"AbuseScope", string(data)}, " ")
}

type AbuseScopeType struct {
	value string
}

type AbuseScopeTypeEnum struct {
	ALL_USERS   AbuseScopeType
	SINGLE_USER AbuseScopeType
}

func GetAbuseScopeTypeEnum() AbuseScopeTypeEnum {
	return AbuseScopeTypeEnum{
		ALL_USERS: AbuseScopeType{
			value: "all_users",
		},
		SINGLE_USER: AbuseScopeType{
			value: "single_user",
		},
	}
}

func (c AbuseScopeType) Value() string {
	return c.value
}

func (c AbuseScopeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AbuseScopeType) UnmarshalJSON(b []byte) error {
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
