package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserGroupInfo struct {

	// 用户组名。
	Name *string `json:"name,omitempty"`

	// 用户组ID。
	Id *string `json:"id,omitempty"`

	// 用户组对应的创建时间，UTC时间：yyyy-MM-ddTHH:mm:ss.SSSZ。
	CreateTime *string `json:"create_time,omitempty"`

	// 用户组描述。
	Description *string `json:"description,omitempty"`

	// 用户列表中用户数。
	UserQuantity *int32 `json:"user_quantity,omitempty"`

	Parent *UserGroupInfo `json:"parent,omitempty"`

	// 用户组域Id。
	RealmId *string `json:"realm_id,omitempty"`

	// 用户组类型。 * AD： AD域用户组 * LOCAL： 本地liteAs用户组
	PlatformType *UserGroupInfoPlatformType `json:"platform_type,omitempty"`

	// 用户组专有名。
	GroupDn *string `json:"group_dn,omitempty"`

	// 用户组域名。
	Domain *string `json:"domain,omitempty"`

	// 用户组sid。
	Sid *string `json:"sid,omitempty"`

	// 用户列表中用户数。
	TotalDesktops *int32 `json:"total_desktops,omitempty"`
}

func (o UserGroupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserGroupInfo struct{}"
	}

	return strings.Join([]string{"UserGroupInfo", string(data)}, " ")
}

type UserGroupInfoPlatformType struct {
	value string
}

type UserGroupInfoPlatformTypeEnum struct {
	AD    UserGroupInfoPlatformType
	LOCAL UserGroupInfoPlatformType
}

func GetUserGroupInfoPlatformTypeEnum() UserGroupInfoPlatformTypeEnum {
	return UserGroupInfoPlatformTypeEnum{
		AD: UserGroupInfoPlatformType{
			value: "AD",
		},
		LOCAL: UserGroupInfoPlatformType{
			value: "LOCAL",
		},
	}
}

func (c UserGroupInfoPlatformType) Value() string {
	return c.value
}

func (c UserGroupInfoPlatformType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserGroupInfoPlatformType) UnmarshalJSON(b []byte) error {
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
