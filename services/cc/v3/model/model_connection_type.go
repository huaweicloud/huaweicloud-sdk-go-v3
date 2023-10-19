package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionType 连接类型。
type ConnectionType struct {
	ConnectionType *ConnectionTypeEnum `json:"connection_type"`
}

func (o ConnectionType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionType struct{}"
	}

	return strings.Join([]string{"ConnectionType", string(data)}, " ")
}
