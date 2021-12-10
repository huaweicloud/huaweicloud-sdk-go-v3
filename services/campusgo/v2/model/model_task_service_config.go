package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业运行时指定的算法配置参数
type TaskServiceConfig struct {
	// 作业运行时指定的具体的算法配置项，以人流检测服务为例

	Common *interface{} `json:"common,omitempty"`
}

func (o TaskServiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskServiceConfig struct{}"
	}

	return strings.Join([]string{"TaskServiceConfig", string(data)}, " ")
}
