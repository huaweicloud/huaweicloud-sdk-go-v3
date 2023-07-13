package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LimitsRequests 资源配置限制
type LimitsRequests struct {

	// cpu核数，大于等于0.01，最大1000；请求不需要带单位
	Cpu *float32 `json:"cpu,omitempty"`

	// 内存大小，单位兆，大于等于0.01，最大1024000。注意：内存的limits值最小为4；请求不需要带单位
	Memory *float32 `json:"memory,omitempty"`

	// Gpu显存大小，单位兆，大于等于0.01，最大1024000；请求不需要带单位
	Gpu *float32 `json:"gpu,omitempty"`

	// Npu个数，大于0，最大1000；请求不需要带单位
	Npu *int32 `json:"npu,omitempty"`
}

func (o LimitsRequests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LimitsRequests struct{}"
	}

	return strings.Join([]string{"LimitsRequests", string(data)}, " ")
}
