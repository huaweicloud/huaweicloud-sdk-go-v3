package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FuncLogConfig 函数绑定日志配置。
type FuncLogConfig struct {

	// 函数绑定日志组名。
	GroupName *string `json:"group_name,omitempty"`

	// 函数绑定日志组ID。
	GroupId *string `json:"group_id,omitempty"`

	// 函数绑定日志流名。
	StreamName *string `json:"stream_name,omitempty"`

	// 函数绑定日志流ID。
	StreamId *string `json:"stream_id,omitempty"`
}

func (o FuncLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FuncLogConfig struct{}"
	}

	return strings.Join([]string{"FuncLogConfig", string(data)}, " ")
}
