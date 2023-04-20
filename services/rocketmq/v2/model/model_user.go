package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type User struct {

	// 用户名。
	AccessKey *string `json:"access_key,omitempty"`

	// 密钥。
	SecretKey *string `json:"secret_key,omitempty"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty"`

	// 默认的主题权限。
	DefaultTopicPerm *UserDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// 默认的消费组权限。
	DefaultGroupPerm *UserDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// 特殊的主题权限。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// 特殊的消费组权限。
	GroupPerms *[]UserGroupPerms `json:"group_perms,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}

type UserDefaultTopicPerm struct {
	value string
}

type UserDefaultTopicPermEnum struct {
	PUB     UserDefaultTopicPerm
	SUB     UserDefaultTopicPerm
	PUB_SUB UserDefaultTopicPerm
	DENY    UserDefaultTopicPerm
}

func GetUserDefaultTopicPermEnum() UserDefaultTopicPermEnum {
	return UserDefaultTopicPermEnum{
		PUB: UserDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UserDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UserDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UserDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UserDefaultTopicPerm) Value() string {
	return c.value
}

func (c UserDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UserDefaultGroupPerm struct {
	value string
}

type UserDefaultGroupPermEnum struct {
	SUB  UserDefaultGroupPerm
	DENY UserDefaultGroupPerm
}

func GetUserDefaultGroupPermEnum() UserDefaultGroupPermEnum {
	return UserDefaultGroupPermEnum{
		SUB: UserDefaultGroupPerm{
			value: "SUB",
		},
		DENY: UserDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UserDefaultGroupPerm) Value() string {
	return c.value
}

func (c UserDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
