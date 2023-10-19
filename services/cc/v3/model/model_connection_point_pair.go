package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionPointPair 带宽包两端接入站点的站点编码，成对配置
type ConnectionPointPair struct {

	// 中心网络连接的两个端点定义，长度固定为2的数组。
	ConnectionPointPair []ConnectionPoint `json:"connection_point_pair"`
}

func (o ConnectionPointPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionPointPair struct{}"
	}

	return strings.Join([]string{"ConnectionPointPair", string(data)}, " ")
}
