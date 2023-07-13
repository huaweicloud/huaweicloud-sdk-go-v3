package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserResponse Response Object
type ShowUserResponse struct {

	// 用户名。
	AccessKey *string `json:"access_key,omitempty"`

	// 密钥。
	SecretKey *string `json:"secret_key,omitempty"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty"`

	// 默认的主题权限。
	DefaultTopicPerm *ShowUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// 默认的消费组权限。
	DefaultGroupPerm *ShowUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// 特殊的主题权限。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// 特殊的消费组权限。
	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty"`
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

type ShowUserResponseDefaultGroupPerm struct {
	value string
}

type ShowUserResponseDefaultGroupPermEnum struct {
	SUB  ShowUserResponseDefaultGroupPerm
	DENY ShowUserResponseDefaultGroupPerm
}

func GetShowUserResponseDefaultGroupPermEnum() ShowUserResponseDefaultGroupPermEnum {
	return ShowUserResponseDefaultGroupPermEnum{
		SUB: ShowUserResponseDefaultGroupPerm{
			value: "SUB",
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
