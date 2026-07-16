package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Health **参数解释：** 健康检查配置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
type Health struct {

	// **参数解释：** 健康检查方式：HTTP 或者 EXEC（命令行）。 **约束限制：** 不涉及。 **取值范围：** - HTTP：超文本传输协议。 - EXEC：命令行。 **默认取值：** 创建的时候，check_method不传值，默认：check_method。 更新的时候，check_method不传值，会报错。
	CheckMethod *string `json:"check_method,omitempty"`

	// **参数解释：** 当健康检查方式为EXEC时必填，配置的命令行。 **约束限制：** 字符长度限制[0, 1024]，不能包含字符：#~^$|%&*<>()'\"[]{} **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Cmd *string `json:"cmd,omitempty"`

	// **参数解释：** 当健康检查方式为HTTP 时必填，配置的请求地址。 **约束限制：** 字符长度限制[0, 1024]，首字符为/，后续字符可以是：字母 数字 中划线 下划线 / : **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 连接协议。默认HTTP。 **约束限制：** 不涉及。 **取值范围：** - HTTPS：超文本传输协议安全版。 - HTTP：超文本传输协议。 - WSS：网络通信协议安全版。 - WS：网络通信协议。 **默认取值：** 不涉及。
	Protocol *string `json:"protocol,omitempty"`

	// **参数解释：** 执行首次探测时，应该等待的时间，默认30秒，最小1秒。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 默认值为30。
	InitialDelaySeconds int32 `json:"initial_delay_seconds"`

	// **参数解释：** 执行探测的超时时间，默认30秒，最小1秒。 **约束限制：** 不涉及。 **取值范围：** 最小值为1秒。 **默认取值：** 默认值为30秒。
	TimeoutSeconds int32 `json:"timeout_seconds"`

	// **参数解释：** 执行健康检查的周期时间，执行探测的频率。默认是10秒，最小1秒。 **约束限制：** 不涉及。 **取值范围：** 最小值为1秒。 **默认取值：** 默认值为10秒。
	PeriodSeconds int32 `json:"period_seconds"`

	// **参数解释：** 探测成功后，至少连续探测失败多少次才被认定为失败。默认是3。最小值是1。 **约束限制：** 不涉及。 **取值范围：** 最小值为1。 **默认取值：** 默认值为3。
	FailureThreshold int32 `json:"failure_threshold"`
}

func (o Health) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Health struct{}"
	}

	return strings.Join([]string{"Health", string(data)}, " ")
}
