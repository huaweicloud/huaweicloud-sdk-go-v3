package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddLogConfigResponse Response Object
type AddLogConfigResponse struct {

	// 添加日志配置返回值
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddLogConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddLogConfigResponse struct{}"
	}

	return strings.Join([]string{"AddLogConfigResponse", string(data)}, " ")
}
