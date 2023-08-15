package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CircuitBreaker 熔断配置，当502/504请求数量或读等待URL请求数量以及占比阈值达到您设置的值时，将触发WAF熔断功能开关，实现宕机保护和读等待URL请求保护
type CircuitBreaker struct {

	// 熔断开关，是否开启连接保护   - true：开启连接保护   - false: 关闭连接保护
	Switch *bool `json:"switch,omitempty"`

	// 502/504数量阈值，每30s累加的502/504数量阈值
	DeadNum *int32 `json:"dead_num,omitempty"`

	// 502/504数量占比(%)，总请求数量中502/504数量占比达到所设定值，并且与数量阈值同时满足时触发宕机保护
	DeadRatio float32 `json:"dead_ratio,omitempty"`

	// 初次触发宕机的保护时间，即WAF将停止转发用户请求的时间。
	BlockTime *int32 `json:"block_time,omitempty"`

	// 连续触发时，保护时间延长最大倍数，叠加周期为3600s。例如，“初次保护时间”设置为180s，“连续触发叠加系数”设置为3。   - 当触发次数为2（即小于3）时，保护时间为360s。   - 当次数大于等于3时，保护时间为540s。   - 当累计保护时间超过1小时（3600s），叠加次数会从头计数。
	SuperpositionNum *int32 `json:"superposition_num,omitempty"`

	// 读等待URL请求数量阈值，读等待URL请求数量到达设定值即触发连接保护
	SuspendNum *int32 `json:"suspend_num,omitempty"`

	// 读等待URL请求数量超过阈值后的熔断时间，达到数量阈值所触发的保护时间，即WAF将停止转发用户请求的时间。
	SusBlockTime *int32 `json:"sus_block_time,omitempty"`
}

func (o CircuitBreaker) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CircuitBreaker struct{}"
	}

	return strings.Join([]string{"CircuitBreaker", string(data)}, " ")
}
