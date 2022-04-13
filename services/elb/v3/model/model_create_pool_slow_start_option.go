package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 慢启动信息。开启慢启动后，将会在设定的时间段（duration）内对新添加到后端服务器组的后端服务器进行预热，转发到该服务器的请求数量线性增加。  当后端服务器组的协议为HTTP/HTTPS时有效，其他协议传入该字段将报错。  [不支持该字段，请勿使用。](tag:dt,dt_test)
type CreatePoolSlowStartOption struct {
	// 慢启动的开关，取值： - true：开启。 - false：关闭，默认值。

	Enable *bool `json:"enable,omitempty"`
	// 慢启动的持续时间，单位：s。默认：30； 取值范围：30~1200

	Duration *int32 `json:"duration,omitempty"`
}

func (o CreatePoolSlowStartOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolSlowStartOption struct{}"
	}

	return strings.Join([]string{"CreatePoolSlowStartOption", string(data)}, " ")
}
