package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUpgradeCmdResponse Response Object
type CreateUpgradeCmdResponse struct {

	// 标准版节点安装/升级命令
	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateUpgradeCmdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUpgradeCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateUpgradeCmdResponse", string(data)}, " ")
}
