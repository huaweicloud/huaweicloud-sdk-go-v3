package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterInstallCmdResponse Response Object
type CreateClusterInstallCmdResponse struct {

	// 标准版节点安装/升级命令
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterInstallCmdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterInstallCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterInstallCmdResponse", string(data)}, " ")
}
