package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskDto 任务请求信息
type TaskDto struct {

	// 任务输入信息，json格式；每个子服务该对象结构不同，框架层不解析具体key，运行态直接透传给子服务REST API处理（参数合法性校验 只能子服务镜像内进行；当前算法镜像提供单独校验接口和处理接口，后续待统一交互机制）
	InputJson *interface{} `json:"input_json"`
}

func (o TaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDto struct{}"
	}

	return strings.Join([]string{"TaskDto", string(data)}, " ")
}
