package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachInstancesUserInfo 分配桌面用户信息。
type AttachInstancesUserInfo struct {

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 桌面所属的用户，当桌面分配成功后此用户可以登录该桌面。只允许输入大写字母、小写字母、数字、中划线（-）和下划线（_）。域类型为LITE_AD时，使用小写字母或者大写字母开头，长度范围为[1-20]。当域类型为LOCAL_AD时，用户名可以使用小写字母或者大写字母或者数字开头，长度范围为[1-32]。Windows桌面用户最长支持20个字符，Linux桌面用户最长支持32个字符。
	UserName *string `json:"user_name,omitempty"`

	// 桌面用户所属的用户组。 - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup *string `json:"user_group,omitempty"`

	// 对象类型，可选值为： - USER：用户。 - GROUP：用户组。
	Type *AttachInstancesUserInfoType `json:"type,omitempty"`
}

func (o AttachInstancesUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesUserInfo struct{}"
	}

	return strings.Join([]string{"AttachInstancesUserInfo", string(data)}, " ")
}

type AttachInstancesUserInfoType struct {
	value string
}

type AttachInstancesUserInfoTypeEnum struct {
	USER  AttachInstancesUserInfoType
	GROUP AttachInstancesUserInfoType
}

func GetAttachInstancesUserInfoTypeEnum() AttachInstancesUserInfoTypeEnum {
	return AttachInstancesUserInfoTypeEnum{
		USER: AttachInstancesUserInfoType{
			value: "USER",
		},
		GROUP: AttachInstancesUserInfoType{
			value: "GROUP",
		},
	}
}

func (c AttachInstancesUserInfoType) Value() string {
	return c.value
}

func (c AttachInstancesUserInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachInstancesUserInfoType) UnmarshalJSON(b []byte) error {
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
