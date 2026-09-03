package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionDetail 连接详情
type ConnectionDetail struct {

	// 客户端IP
	ClientIp *string `json:"client_ip,omitempty"`

	// 连接数
	Count *int32 `json:"count,omitempty"`
}

func (o ConnectionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionDetail struct{}"
	}

	return strings.Join([]string{"ConnectionDetail", string(data)}, " ")
}
