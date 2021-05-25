package model

import (
	"encoding/json"

	"strings"
)

// 云服务器元数据
type VmMetaData struct {
	// Windows弹性云服务器Administrator用户的密码。

	AdminPass *string `json:"admin_pass,omitempty"`
}

func (o VmMetaData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VmMetaData struct{}"
	}

	return strings.Join([]string{"VmMetaData", string(data)}, " ")
}
