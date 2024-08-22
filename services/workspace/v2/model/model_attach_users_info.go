package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachUsersInfo 桌面分配用户信息。
type AttachUsersInfo struct {

	// 当type字段为USER时，填写用户id；当type字段为GROUP时，填写用户组id，后端服务会校验组id是否存在；
	Id *string `json:"id,omitempty"`

	// 桌面分配对象的名称，当type类型USER时填写用户名字；当type类型GROUP时填写用户组名称。 - 当type类型USER时：桌面所属的用户，当桌面分配成功后此用户可以登录该桌面。只允许输入大写字母、小写字母、数字、中划线（-）和下划线（_）。域类型为LITE_AD时，使用小写字母或者大写字母开头，长度范围为[1-20]。当域类型为LOCAL_AD时，用户名可以使用小写字母或者大写字母或者数字开头，长度范围为[1-64]。Windows桌面用户最长支持20个字符，Linux桌面用户最长支持64个字符。后端服务会校验用户名是否存在，并且用户名不能与机器名重复。 - 当type类型GROUP时：只能为中文、字母、数字及特殊符号-_。
	Name *string `json:"name,omitempty"`

	// 桌面用户所属的用户组。 - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup *string `json:"user_group,omitempty"`

	// 对象类型，可选值为： - USER：用户。 - GROUP：用户组。
	Type *string `json:"type,omitempty"`
}

func (o AttachUsersInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachUsersInfo struct{}"
	}

	return strings.Join([]string{"AttachUsersInfo", string(data)}, " ")
}
