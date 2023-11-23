package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstancesSessionStatisticsRespondBodyTopSourceIps struct {

	// 客户端ip地址。
	ClientIp *string `json:"client_ip,omitempty"`

	// 客户端连接数。
	ConnectionCount *int32 `json:"connection_count,omitempty"`
}

func (o ListInstancesSessionStatisticsRespondBodyTopSourceIps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesSessionStatisticsRespondBodyTopSourceIps struct{}"
	}

	return strings.Join([]string{"ListInstancesSessionStatisticsRespondBodyTopSourceIps", string(data)}, " ")
}
