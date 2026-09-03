package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDesktopPoolUserReq 桌面池绑定用户请求。
type AttachDesktopPoolUserReq struct {

	// 用户名。
	Username string `json:"username"`

	// 域名。
	DomainName *string `json:"domain_name,omitempty"`
}

func (o AttachDesktopPoolUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDesktopPoolUserReq struct{}"
	}

	return strings.Join([]string{"AttachDesktopPoolUserReq", string(data)}, " ")
}
