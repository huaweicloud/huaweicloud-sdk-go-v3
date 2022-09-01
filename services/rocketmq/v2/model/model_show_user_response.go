package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowUserResponse struct {

	// 用户名。
	AccessKey *string `json:"access_key,omitempty" xml:"access_key"`

	// 密钥。
	SecretKey *string `json:"secret_key,omitempty" xml:"secret_key"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty" xml:"white_remote_address"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty" xml:"admin"`

	// 默认的主题权限。
	DefaultTopicPerm *ShowUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty" xml:"default_topic_perm"`

	// 默认的消费组权限。
	DefaultGroupPerm *ShowUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty" xml:"default_group_perm"`

	// 特殊的主题权限。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty" xml:"topic_perms"`

	// 特殊的消费组权限。
	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty" xml:"group_perms"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserResponse struct{}"
	}

	return strings.Join([]string{"ShowUserResponse", string(data)}, " ")
}

type ShowUserResponseDefaultTopicPerm struct {
	value string
}

type ShowUserResponseDefaultTopicPermEnum struct {
	PUB     ShowUserResponseDefaultTopicPerm
	SUB     ShowUserResponseDefaultTopicPerm
	PUB_SUB ShowUserResponseDefaultTopicPerm
	DENY    ShowUserResponseDefaultTopicPerm
}

func GetShowUserResponseDefaultTopicPermEnum() ShowUserResponseDefaultTopicPermEnum {
	return ShowUserResponseDefaultTopicPermEnum{
		PUB: ShowUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: ShowUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: ShowUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: ShowUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c ShowUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c ShowUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type ShowUserResponseDefaultGroupPerm struct {
	value string
}

type ShowUserResponseDefaultGroupPermEnum struct {
	PUB     ShowUserResponseDefaultGroupPerm
	SUB     ShowUserResponseDefaultGroupPerm
	PUB_SUB ShowUserResponseDefaultGroupPerm
	DENY    ShowUserResponseDefaultGroupPerm
}

func GetShowUserResponseDefaultGroupPermEnum() ShowUserResponseDefaultGroupPermEnum {
	return ShowUserResponseDefaultGroupPermEnum{
		PUB: ShowUserResponseDefaultGroupPerm{
			value: "PUB",
		},
		SUB: ShowUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		PUB_SUB: ShowUserResponseDefaultGroupPerm{
			value: "PUB|SUB",
		},
		DENY: ShowUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c ShowUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c ShowUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
