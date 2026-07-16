package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JupyterLab JupyterLab连接信息。
type JupyterLab struct {

	// 训练作业的JupyterLab地址。
	Url *string `json:"url,omitempty"`

	// 训练作业的JupyterLab token。
	Token *string `json:"token,omitempty"`
}

func (o JupyterLab) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JupyterLab struct{}"
	}

	return strings.Join([]string{"JupyterLab", string(data)}, " ")
}
