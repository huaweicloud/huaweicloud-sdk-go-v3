package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemUserName 系统用户名
type SystemUserName struct {
}

func (o SystemUserName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemUserName struct{}"
	}

	return strings.Join([]string{"SystemUserName", string(data)}, " ")
}
