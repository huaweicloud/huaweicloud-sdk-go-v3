package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VCN服务器信息，只有输入为vcn时才可以且必须使用
type TaskInputVcn struct {

	// VCN服务器的IP地址
	Ip string `json:"ip"`

	// VCN服务器的端口号
	Port int32 `json:"port"`

	// VCN服务器的账号名
	Username string `json:"username"`

	// VCN服务器的与账号对应的密码
	Password string `json:"password"`
}

func (o TaskInputVcn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInputVcn struct{}"
	}

	return strings.Join([]string{"TaskInputVcn", string(data)}, " ")
}
