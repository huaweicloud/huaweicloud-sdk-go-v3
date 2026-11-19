package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerlessScaleToZeroPolicy 设置Serverless自动启停策略请求体。
type UpdateServerlessScaleToZeroPolicy struct {

	// **参数解释**：  是否开启Scale to Zero自动启停。  **约束限制**：  仅Serverless实例支持。  **取值范围**：  - true：开启。实例空闲达idle_duration_seconds阈值时算力自动缩容到零（停止VM计费），有新连接时自动唤醒恢复。 - false：关闭。实例正常计费。  **默认取值**：  false。
	ScaleToZero bool `json:"scale_to_zero"`

	// **参数解释**：  触发缩容到零的空闲检测时长（秒）。实例所有节点无活跃用户连接达到该时长后，自动缩容到零。  **约束限制**：  - 当scale_to_zero为true时，该参数必填。 - 当scale_to_zero为false时，该参数可不传。  **取值范围**：  300～3600。  **默认取值**：  不涉及。
	IdleDurationSeconds *int32 `json:"idle_duration_seconds,omitempty"`
}

func (o UpdateServerlessScaleToZeroPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerlessScaleToZeroPolicy struct{}"
	}

	return strings.Join([]string{"UpdateServerlessScaleToZeroPolicy", string(data)}, " ")
}
