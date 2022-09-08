package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标准版节点获取安装命令请求体
type CreateInstallCmdRequestDto struct {

	// 安装命令执行的主机标签DEFAULT|MASTER|SLAVE
	HostTag *string `json:"host_tag,omitempty"`
}

func (o CreateInstallCmdRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstallCmdRequestDto struct{}"
	}

	return strings.Join([]string{"CreateInstallCmdRequestDto", string(data)}, " ")
}
