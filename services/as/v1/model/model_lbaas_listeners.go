package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 增强型负载均衡器
type LbaasListeners struct {
	// 后端云服务器组ID

	PoolId string `json:"pool_id"`
	// 后端协议号，指后端云服务器监听的端口，取值范围[1,65535]

	ProtocolPort int32 `json:"protocol_port"`
	// 权重，指后端云服务器经分发得到的请求数量的比例，取值范围[0, 100]。

	Weight int32 `json:"weight"`
}

func (o LbaasListeners) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbaasListeners struct{}"
	}

	return strings.Join([]string{"LbaasListeners", string(data)}, " ")
}
