package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AuthorizedObjects 查询桌面池授权的用户、用户组响应的用户/用户组信息
type AuthorizedObjects struct {

	// 绑定对象类型枚举。  - USER：用户 - USER_GROUP：用户组
	ObjectType AuthorizedObjectsObjectType `json:"object_type"`

	// 用户/用户组id
	ObjectId string `json:"object_id"`

	// 用户/用户组名称
	ObjectName string `json:"object_name"`

	// 桌面用户所属的用户权限组。  - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup string `json:"user_group"`

	// 创建时间。格式为：UTC格式，例如“2022-05-11T11:45:42.000Z”。
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o AuthorizedObjects) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizedObjects struct{}"
	}

	return strings.Join([]string{"AuthorizedObjects", string(data)}, " ")
}

type AuthorizedObjectsObjectType struct {
	value string
}

type AuthorizedObjectsObjectTypeEnum struct {
	USER       AuthorizedObjectsObjectType
	USER_GROUP AuthorizedObjectsObjectType
}

func GetAuthorizedObjectsObjectTypeEnum() AuthorizedObjectsObjectTypeEnum {
	return AuthorizedObjectsObjectTypeEnum{
		USER: AuthorizedObjectsObjectType{
			value: "USER",
		},
		USER_GROUP: AuthorizedObjectsObjectType{
			value: "USER_GROUP",
		},
	}
}

func (c AuthorizedObjectsObjectType) Value() string {
	return c.value
}

func (c AuthorizedObjectsObjectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AuthorizedObjectsObjectType) UnmarshalJSON(b []byte) error {
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
