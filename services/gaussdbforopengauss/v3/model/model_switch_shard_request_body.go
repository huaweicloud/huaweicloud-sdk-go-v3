package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchShardRequestBody struct {

	// 节点列表，支持对单个或者多个DN分片做主备切换。节点信息为将要升主的备DN分片对应的节点id和组件id。
	Shards []Shards `json:"shards" xml:"shards"`
}

func (o SwitchShardRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchShardRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchShardRequestBody", string(data)}, " ")
}
