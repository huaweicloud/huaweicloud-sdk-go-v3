package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoliciesPeripheralsDriverInterfaceRedirection 驱动接口重定向。
type PoliciesPeripheralsDriverInterfaceRedirection struct {

	// 自定义驱动列表。（填写安装在终端的驱动文件名或驱动文件的全路径，支持配置多个，多个之间以\";\"隔开），长度不能超过1000个字符。
	ApiRedirDriverList *string `json:"api_redir_driver_list,omitempty"`
}

func (o PoliciesPeripheralsDriverInterfaceRedirection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesPeripheralsDriverInterfaceRedirection struct{}"
	}

	return strings.Join([]string{"PoliciesPeripheralsDriverInterfaceRedirection", string(data)}, " ")
}
