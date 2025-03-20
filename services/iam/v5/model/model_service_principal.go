package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServicePrincipal 服务主体，由\"service.\"开头，后跟一个长度为1到56个字符，只包含字母、数字和\"-\"的字符串。
type ServicePrincipal struct {
}

func (o ServicePrincipal) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServicePrincipal struct{}"
	}

	return strings.Join([]string{"ServicePrincipal", string(data)}, " ")
}
