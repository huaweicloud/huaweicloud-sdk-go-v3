package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserTopicPerms struct {

	// 主题名称。
	Name *string `json:"name,omitempty"`

	// 权限。
	Perm *UserTopicPermsPerm `json:"perm,omitempty"`
}

func (o UserTopicPerms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserTopicPerms struct{}"
	}

	return strings.Join([]string{"UserTopicPerms", string(data)}, " ")
}

type UserTopicPermsPerm struct {
	value string
}

type UserTopicPermsPermEnum struct {
	PUB     UserTopicPermsPerm
	SUB     UserTopicPermsPerm
	PUB_SUB UserTopicPermsPerm
	DENY    UserTopicPermsPerm
}

func GetUserTopicPermsPermEnum() UserTopicPermsPermEnum {
	return UserTopicPermsPermEnum{
		PUB: UserTopicPermsPerm{
			value: "PUB",
		},
		SUB: UserTopicPermsPerm{
			value: "SUB",
		},
		PUB_SUB: UserTopicPermsPerm{
			value: "PUB|SUB",
		},
		DENY: UserTopicPermsPerm{
			value: "DENY",
		},
	}
}

func (c UserTopicPermsPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserTopicPermsPerm) UnmarshalJSON(b []byte) error {
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
