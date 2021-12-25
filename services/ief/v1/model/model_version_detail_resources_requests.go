package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 容器需要使用的最小资源
type VersionDetailResourcesRequests struct {
	// cpu核数，大于等于0.01，最大1000；请求不需要带单位

	Cpu *string `json:"cpu,omitempty"`
	// 内存大小，单位兆，大于等于0.01，最大1024000。注意：内存的limits值最小为4；请求不需要带单位

	Memory *string `json:"memory,omitempty"`
	// Gpu显存大小，单位兆，大于等于0.01，最大1024000；请求不需要带单位

	Gpu *string `json:"gpu,omitempty"`
	// Npu个数，大于0，最大1000；请求不需要带单位

	Npu *string `json:"npu,omitempty"`
}

func (o VersionDetailResourcesRequests) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetailResourcesRequests struct{}"
	}

	return strings.Join([]string{"VersionDetailResourcesRequests", string(data)}, " ")
}
