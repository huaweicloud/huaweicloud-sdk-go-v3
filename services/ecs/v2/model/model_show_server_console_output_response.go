package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerConsoleOutputResponse Response Object
type ShowServerConsoleOutputResponse struct {

	// 云服务器控制台日志。
	Output         *string `json:"output,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowServerConsoleOutputResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerConsoleOutputResponse struct{}"
	}

	return strings.Join([]string{"ShowServerConsoleOutputResponse", string(data)}, " ")
}
