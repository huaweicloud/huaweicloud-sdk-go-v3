package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePoolSlowStartOption 参数解释：慢启动信息。开启慢启动后，将会在设定的时间段（duration）内对新添加到后端服务器组的后端服务器进行预热，转发到该服务器的请求数量线性增加。  约束限制： - 当后端服务器组的协议为HTTP/HTTPS/GRPC时有效，其他协议传入该字段将报错。 - 慢启动与会话保持不能同时开启。若都开启则会导致会话保持失效。  [网关型LB，不支持该特性，请勿使用。](tag:hws_eu) [荷兰region不支持该字段，请勿使用。](tag:dt)
type CreatePoolSlowStartOption struct {

	// 参数解释：慢启动的开关。  取值范围： - true：开启。 - false：关闭。  默认取值：false。
	Enable *bool `json:"enable,omitempty"`

	// 参数解释：慢启动的持续时间。  取值范围：30-1200，单位：秒。  默认取值：30。
	Duration *int32 `json:"duration,omitempty"`
}

func (o CreatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSlowStartOption", string(data)}, " ")
}
