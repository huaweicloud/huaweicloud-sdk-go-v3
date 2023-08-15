package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommandInfo struct {

	// 原高危命令，当前支持的有：keys、flushdb、flushall、hgetall、hkeys、hvals、smembers
	OriginName string `json:"origin_name"`

	// 当前生效的命令名称，当为空时表示命令禁用，允许修改为30个字符以内数字、字母和下划线的组合
	Name string `json:"name"`
}

func (o CommandInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommandInfo struct{}"
	}

	return strings.Join([]string{"CommandInfo", string(data)}, " ")
}
