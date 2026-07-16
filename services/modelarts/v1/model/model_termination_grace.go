package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TerminationGrace **参数解释：** 开启后，支持设置停机时间及停机命令等，避免正在处理的请求被强制中断，从而提高系统的可用性和稳定性。 **约束限制：** 长度不超过255。 **取值范围：** - 协议范围：http/https。 - 端口范围：1-65535。 - 地址范围：仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径，非斜杠（/）开头。 **默认取值：** 不涉及。
type TerminationGrace struct {

	// **参数解释：** 停机命令在容器收到停止信号时触发，但必须在停机时间的宽限期内完成，否则容器会被强制停止。您可以通过该命令精细化操作，如关闭数据库连接、释放文件句柄、停止子进程等。 **约束限制：** 长度不超过255。 **取值范围：** - 协议范围：http/https。 - 端口范围：1-65535。 - 地址范围：仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径，非斜杠（/）开头。 **默认取值：** 不涉及。
	PreStopCmd *string `json:"pre_stop_cmd,omitempty"`

	// **参数解释：** 该参数表示 Pod 收到停止信号到强制停止的最大时间窗口，用于 Pod 执行清理操作（如关闭连接、释放资源、保存状态等）。 **约束限制：** 长度不超过255。 **取值范围：** - 协议范围：http/https。 - 端口范围：1-65535。 - 地址范围：仅包含字母、数字、点号（.）、中划线（-)、下划线（_）、斜杠（/）的路径，非斜杠（/）开头。 **默认取值：** 不涉及。
	TerminationGracePeriodSeconds *int32 `json:"termination_grace_period_seconds,omitempty"`
}

func (o TerminationGrace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TerminationGrace struct{}"
	}

	return strings.Join([]string{"TerminationGrace", string(data)}, " ")
}
