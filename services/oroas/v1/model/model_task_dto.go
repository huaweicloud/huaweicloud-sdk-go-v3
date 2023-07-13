package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskDto 任务请求信息
type TaskDto struct {

	// 任务输入信息，json格式；每个子服务该对象结构不同，框架层不解析具体key，运行态直接透传给子服务REST API处理（参数合法性校验 只能子服务镜像内进行；当前算法镜像提供单独校验接口和处理接口，后续待统一交互机制）
	InputJson *interface{} `json:"input_json"`

	// 任务输入信息为文件格式，传入值为租户OBS对应的文件绝对路径；在请求信息大于12MB情形使用该参数，需用户在Console进行OBS委托授权方可使用（和input_json二选一），暂不开放该功能
	InputUrl *string `json:"input_url,omitempty"`

	// 任务输出信息为文件格式，传入值为租户OBS对应的待存储路径前缀（和input_url成对使用），文件名服务端固定用task_id命名；在响应信息大于12MB情形使用该参数，需用户在Console进行OBS委托授权方可使用，暂不开放该功能
	OutputUrl *string `json:"output_url,omitempty"`
}

func (o TaskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDto struct{}"
	}

	return strings.Join([]string{"TaskDto", string(data)}, " ")
}
