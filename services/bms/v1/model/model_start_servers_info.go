package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// os-start字段数据结构说明
type StartServersInfo struct {
	// 裸金属服务器ID列表，详情请参见表3 servers字段数据结构说明

	Servers []ServersList `json:"servers"`
}

func (o StartServersInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartServersInfo struct{}"
	}

	return strings.Join([]string{"StartServersInfo", string(data)}, " ")
}
