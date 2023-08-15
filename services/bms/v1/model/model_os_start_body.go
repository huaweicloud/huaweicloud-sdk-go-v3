package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsStartBody 启动裸金属服务器接口请求结构体
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
