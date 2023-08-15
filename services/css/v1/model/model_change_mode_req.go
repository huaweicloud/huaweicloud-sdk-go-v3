package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeModeReq struct {

	// 是否开启安全模式。 - true: 开启安全模式。 - false: 关闭安全模式。 默认为：true。
	AuthorityEnable bool `json:"authorityEnable"`

	// 安全模式下集群密码。
	AdminPwd *string `json:"adminPwd,omitempty"`

	// 是否开启HTTPS。 - true: 开启HTTPS。 - false: 关闭HTTPS。 默认为：true。
	HttpsEnable bool `json:"httpsEnable"`
}

func (o ChangeModeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeModeReq struct{}"
	}

	return strings.Join([]string{"ChangeModeReq", string(data)}, " ")
}
