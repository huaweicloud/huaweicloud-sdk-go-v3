package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Desktop struct {

	// 桌面所属的用户，当桌面创建成功后此用户可以登录该桌面。只允许输入大写字母、小写字母、数字、中划线（-）和下划线（_）。域类型为LITE_AD时，使用小写字母或者大写字母开头，长度范围为[1-20]。当域类型为LOCAL_AD时，用户名可以使用小写字母或者大写字母或者数字开头，长度范围为[1-32]。Windows桌面用户最长支持20个字符，Linux桌面用户最长支持32个字符。
	UserName string `json:"user_name"`

	// 合法用户邮箱，桌面创建成功后系统会通过发送邮件的方式通知用户。
	UserEmail *string `json:"user_email,omitempty"`

	// 合法用户手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 桌面用户所属的用户组。  - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup *string `json:"user_group,omitempty"`

	// 桌面名，桌面名必须保证唯一。桌面名称只允许输入大写字母、小写字母、数字、中划线，以字母或数字开头、不能以中划线结尾，长度范围为1~15。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面名称前缀，不指定\"computer_name\"时生效。
	DesktopNamePrefix *string `json:"desktop_name_prefix,omitempty"`
}

func (o Desktop) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Desktop struct{}"
	}

	return strings.Join([]string{"Desktop", string(data)}, " ")
}
