package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionDrain 参数解释：后端服务器的延迟注销的功能配置（只针对TCP、UDP、QUIC协议类型的后端服务器组和TCP、UDP协议类的监听器）。 以下场景会触发： - 服务器从后端服务器组中移除 - 后端云服务健康检查状态异常 - 后端服务器权重修改为0
type ConnectionDrain struct {

	// 参数解释：延迟注销功能开关。  取值范围：true 开启，false 关闭。  默认取值：true。
	Enable *bool `json:"enable,omitempty"`

	// 参数解释：延迟注销时间。  取值范围：10~4000，单位：秒。
	Timeout *int32 `json:"timeout,omitempty"`
}

func (o ConnectionDrain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionDrain struct{}"
	}

	return strings.Join([]string{"ConnectionDrain", string(data)}, " ")
}
