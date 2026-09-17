package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateReinstallCmdResponse Response Object
type CreateReinstallCmdResponse struct {

	// 标准版节点安装/升级命令
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateReinstallCmdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateReinstallCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateReinstallCmdResponse", string(data)}, " ")
}
