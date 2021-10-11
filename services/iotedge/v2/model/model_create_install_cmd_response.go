package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateInstallCmdResponse struct {
	// 标准版节点安装命令

	Cmd            *string `json:"cmd,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstallCmdResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstallCmdResponse struct{}"
	}

	return strings.Join([]string{"CreateInstallCmdResponse", string(data)}, " ")
}
