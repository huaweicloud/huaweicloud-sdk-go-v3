package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例组信息。
type GroupResponseItem struct {

	// 节点类型。 取值： - shard - config - mongos - replica - single
	Type string `json:"type" xml:"type"`

	// 组ID。节点类型为shard和config时，该参数有效。
	Id string `json:"id" xml:"id"`

	// 组名称。节点类型为shard和config时，该参数有效。
	Name string `json:"name" xml:"name"`

	// 组状态。节点类型为shard和config时，该参数有效。
	Status string `json:"status" xml:"status"`

	Volume *Volume `json:"volume" xml:"volume"`

	// 节点信息。
	Nodes []NodeItem `json:"nodes" xml:"nodes"`
}

func (o GroupResponseItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupResponseItem struct{}"
	}

	return strings.Join([]string{"GroupResponseItem", string(data)}, " ")
}
