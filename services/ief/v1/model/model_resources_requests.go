package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器需要使用的最小资源
type ResourcesRequests struct {
	// cpu核数，大于等于0.01，最大1000；请求不需要带单位

	Cpu *float32 `json:"cpu,omitempty"`
	// 内存大小，单位兆，大于等于0.01，最大1024000。注意：内存的limits值最小为4；请求不需要带单位

	Memory *float32 `json:"memory,omitempty"`
	// Gpu显存大小，单位兆，大于等于0.01，最大1024000；请求不需要带单位

	Gpu *float32 `json:"gpu,omitempty"`
	// Npu个数，大于0，最大1000；请求不需要带单位

	Npu *float32 `json:"npu,omitempty"`
}

func (o ResourcesRequests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesRequests struct{}"
	}

	return strings.Join([]string{"ResourcesRequests", string(data)}, " ")
}
