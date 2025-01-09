package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerGroupStatus 服务器组状态统计。
type ServerGroupStatus struct {

	// 对应状态的服务器数量，参考ServerStatus。
	ApsStatus map[string]int32 `json:"aps_status,omitempty"`
}

func (o ServerGroupStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerGroupStatus struct{}"
	}

	return strings.Join([]string{"ServerGroupStatus", string(data)}, " ")
}
