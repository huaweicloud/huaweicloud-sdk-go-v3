package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业运行时指定的算法配置参数，见园区智能体API参考文档[API参考文档](https://support.huaweicloud.com/api-campusgo/campusgo_03_0013.html)
type TaskServiceConfig struct {
	// 根据不同服务，填写对应的服务配置参数json结构体

	Common *interface{} `json:"common,omitempty"`
}

func (o TaskServiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskServiceConfig struct{}"
	}

	return strings.Join([]string{"TaskServiceConfig", string(data)}, " ")
}
