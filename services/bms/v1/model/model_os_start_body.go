package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动裸金属服务器接口请求结构体
type OsStartBody struct {
	OsStart *StartServersInfo `json:"os-start"`
}

func (o OsStartBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsStartBody struct{}"
	}

	return strings.Join([]string{"OsStartBody", string(data)}, " ")
}
