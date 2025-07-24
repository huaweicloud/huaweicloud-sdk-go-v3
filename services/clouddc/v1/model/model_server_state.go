package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerState 服务器运行状态对象
type ServerState struct {

	// 数量
	Count int32 `json:"count"`

	// 数据中心名称对应数量映射
	DcStats map[string]int32 `json:"dc_stats"`
}

func (o ServerState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerState struct{}"
	}

	return strings.Join([]string{"ServerState", string(data)}, " ")
}
