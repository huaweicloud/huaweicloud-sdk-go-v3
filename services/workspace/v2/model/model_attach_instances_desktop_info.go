package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachInstancesDesktopInfo 分配用户请求桌面信息。
type AttachInstancesDesktopInfo struct {

	// 待分配的桌面ID。
	DesktopId string `json:"desktop_id"`

	// 桌面所属的用户，当桌面分配成功后此用户可以登录该桌面。只允许输入大写字母、小写字母、数字、中划线（-）和下划线（_）。域类型为LITE_AD时，使用小写字母或者大写字母开头，长度范围为[1-20]。当域类型为LOCAL_AD时，用户名可以使用小写字母或者大写字母或者数字开头，长度范围为[1-32]。Windows桌面用户最长支持20个字符，Linux桌面用户最长支持32个字符。
	UserName *string `json:"user_name,omitempty"`

	// 合法用户邮箱，桌面分配成功后系统会通过发送邮件的方式通知用户
	UserEmail *string `json:"user_email,omitempty"`

	// 桌面用户所属的用户组。  - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup *string `json:"user_group,omitempty"`

	// 桌面名，桌面名必须保证唯一。只允许输入大写字母、小写字母、数字、中划线（-）和下划线（_）；以字母开头且不允许以中划线（-）结尾；长度范围为[1-15]。
	ComputerName *string `json:"computer_name,omitempty"`

	// 该字段只有当解绑和绑定为同一个用户时生效。表示绑定时是否清理桌面数据，true：清理，false：不清理，默认值为true。
	IsClearData *bool `json:"is_clear_data,omitempty"`

	// 待分配的用户信息列表。
	AttachUserInfos *[]AttachInstancesUserInfo `json:"attach_user_infos,omitempty"`
}

func (o AttachInstancesDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachInstancesDesktopInfo struct{}"
	}

	return strings.Join([]string{"AttachInstancesDesktopInfo", string(data)}, " ")
}
