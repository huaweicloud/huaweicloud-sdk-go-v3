package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceProcessInfo 进程信息
type CancelInstanceProcessInfo struct {

	// 进程ID
	Id int64 `json:"id"`

	// 用户名
	User string `json:"user"`
}

func (o CancelInstanceProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceProcessInfo struct{}"
	}

	return strings.Join([]string{"CancelInstanceProcessInfo", string(data)}, " ")
}
