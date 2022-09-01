package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业运行时指定的算法配置参数，部分服务需填且必填。
type TaskServiceConfig struct {

	// 作业运行时指定的具体的算法配置项，部分服务需填且必填。整体呈json格式，具体配置项参见相应算法服务的说明。
	Common *interface{} `json:"common,omitempty" xml:"common"`
}

func (o TaskServiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskServiceConfig struct{}"
	}

	return strings.Join([]string{"TaskServiceConfig", string(data)}, " ")
}
