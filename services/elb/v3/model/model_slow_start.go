package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SlowStart 慢启动信息。开启慢启动后，将会在设定的时间段（duration）内对新添加到后端服务器组的后端服务器进行预热，转发到该服务器的请求数量线性增加。  当后端服务器组的协议为HTTP/HTTPS时有效，其他协议传入该字段将报错。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
type SlowStart struct {

	// 慢启动的开关，默认值：false； true：开启； false：关闭
	Enable bool `json:"enable"`

	// 慢启动的持续时间。取值：30~1200s，默认30s；
	Duration int32 `json:"duration"`
}

func (o SlowStart) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowStart struct{}"
	}

	return strings.Join([]string{"SlowStart", string(data)}, " ")
}
