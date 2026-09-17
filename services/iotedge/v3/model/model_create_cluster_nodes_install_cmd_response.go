package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterNodesInstallCmdResponse Response Object
type CreateClusterNodesInstallCmdResponse struct {

	// 标准版节点安装/升级命令
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterNodesInstallCmdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterNodesInstallCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterNodesInstallCmdResponse", string(data)}, " ")
}
