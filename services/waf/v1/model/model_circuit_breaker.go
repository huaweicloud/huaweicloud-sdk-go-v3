package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 熔断配置
type CircuitBreaker struct {

	// 熔断开关
	Switch *bool `json:"switch,omitempty"`

	// 源站不可达数量阈值
	DeadNum *int32 `json:"dead_num,omitempty"`

	// 源站不可达比例阈值
	DeadRatio float32 `json:"dead_ratio,omitempty"`

	// 源站不可达熔断时间
	BlockTime *int32 `json:"block_time,omitempty"`

	// 熔断阈值叠加次数
	SuperpositionNum *int32 `json:"superposition_num,omitempty"`

	// 连接数占用阈值
	SuspendNum *int32 `json:"suspend_num,omitempty"`

	// 连接数占用熔断时间
	SusBlockTime *int32 `json:"sus_block_time,omitempty"`
}

func (o CircuitBreaker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CircuitBreaker struct{}"
	}

	return strings.Join([]string{"CircuitBreaker", string(data)}, " ")
}
