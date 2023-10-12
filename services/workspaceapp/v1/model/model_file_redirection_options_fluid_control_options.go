package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileRedirectionOptionsFluidControlOptions 流控控制项。
type FileRedirectionOptionsFluidControlOptions struct {

	// 网络优的延时阈值（ms）。取值范围为[1-1000]。默认：30。
	GoodNetworkLatency *int32 `json:"good_network_latency,omitempty"`

	// 网络一般的延时阈值（ms）。取值范围为[1-1000]。默认：70。
	NormalNetworkLatency *int32 `json:"normal_network_latency,omitempty"`

	// 网络差的延时阈值（ms）。取值范围为[1-1000]。默认：100。
	PoorNetworkLatency *int32 `json:"poor_network_latency,omitempty"`

	// 降速步伐（KB）。取值范围为[1-100]。默认：20。
	ReducingStep *int32 `json:"reducing_step,omitempty"`

	// 慢增速步伐（KB）。取值范围为[1-100]。默认：10。
	SlowIncreasingStep *int32 `json:"slow_increasing_step,omitempty"`

	// 快增速步伐（KB）。取值范围为[1-100]。默认：20。
	QuickIncreasingStep *int32 `json:"quick_increasing_step,omitempty"`

	// 传输初始速度（KB/s）。取值范围为[1-10240]。默认：1024。
	StartSpeed *int32 `json:"start_speed,omitempty"`

	// 测速块大小（KB）。取值范围为[64-1024]。默认：64。
	TestBlockSize *int32 `json:"test_block_size,omitempty"`

	// 测速块时间间隔（ms）。取值范围为[1000-100000]。默认：10000。
	TestTimeGap *int32 `json:"test_time_gap,omitempty"`
}

func (o FileRedirectionOptionsFluidControlOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionOptionsFluidControlOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionOptionsFluidControlOptions", string(data)}, " ")
}
