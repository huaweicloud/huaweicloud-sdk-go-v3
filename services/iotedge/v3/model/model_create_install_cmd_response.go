package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInstallCmdResponse struct {

	// 标准版节点安装/升级命令
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstallCmdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstallCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateInstallCmdResponse", string(data)}, " ")
}
