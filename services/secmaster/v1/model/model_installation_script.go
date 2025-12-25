package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstallationScript struct {

	// 方法
	Commands *string `json:"commands,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`
}

func (o InstallationScript) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallationScript struct{}"
	}

	return strings.Join([]string{"InstallationScript", string(data)}, " ")
}
