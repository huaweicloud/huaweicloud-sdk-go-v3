package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSerialConsoleOptionsOption
type UpdateSerialConsoleOptionsOption struct {

	// 弹性云服务器云主机是否支持串口登录
	Enabled *bool `json:"enabled,omitempty"`
}

func (o UpdateSerialConsoleOptionsOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSerialConsoleOptionsOption struct{}"
	}

	return strings.Join([]string{"UpdateSerialConsoleOptionsOption", string(data)}, " ")
}
