package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeModeReq struct {

	// 是否开启安全模式。 - true: 开启安全模式。 - false: 关闭安全模式。 默认为：true。
	AuthorityEnable bool `json:"authority_enable"`

	// 安全模式下集群密码。
	AdminPwd *string `json:"admin_pwd,omitempty"`

	// 是否开启HTTPS。 - true: 开启HTTPS。 - false: 关闭HTTPS。 默认为：true。
	HttpsEnable bool `json:"https_enable"`
}

func (o ChangeModeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeModeReq struct{}"
	}

	return strings.Join([]string{"ChangeModeReq", string(data)}, " ")
}
