package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsReinstallBody 重装裸金属服务器操作系统接口请求结构体
type OsReinstallBody struct {
	OsReinstall *OsReinstall `json:"os-reinstall"`
}

func (o OsReinstallBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsReinstallBody struct{}"
	}

	return strings.Join([]string{"OsReinstallBody", string(data)}, " ")
}
