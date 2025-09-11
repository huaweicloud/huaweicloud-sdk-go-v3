package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerList 服务列表
type ServerList struct {

	// server列表
	Server *[]ResponseServer `json:"server,omitempty"`

	// 数量
	Total *int32 `json:"total,omitempty"`
}

func (o ServerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerList struct{}"
	}

	return strings.Join([]string{"ServerList", string(data)}, " ")
}
