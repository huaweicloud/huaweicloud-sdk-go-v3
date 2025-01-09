package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessControl struct {

	// IP接入控制。IP接入控制需输入有效的IP地址和子网掩码，IP和子网掩码间以\"|\"拼接组成一组，如果有多组用\";\"分隔。如：IP|掩码;IP|掩码;IP|掩码。
	IpAccessControl *string `json:"ip_access_control,omitempty"`
}

func (o AccessControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControl struct{}"
	}

	return strings.Join([]string{"AccessControl", string(data)}, " ")
}
