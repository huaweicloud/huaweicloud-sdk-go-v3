package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RespImage 弹性云服务器镜像信息。
type RespImage struct {

	// 弹性云服务器镜像ID。
	Id *string `json:"id,omitempty"`
}

func (o RespImage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespImage struct{}"
	}

	return strings.Join([]string{"RespImage", string(data)}, " ")
}
