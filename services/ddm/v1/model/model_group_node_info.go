package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupNodeInfo struct {

	// 节点ID。
	Id string `json:"id"`

	// 节点名称。
	Name string `json:"name"`

	// 节点所在AZ。
	Az string `json:"az"`
}

func (o GroupNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupNodeInfo struct{}"
	}

	return strings.Join([]string{"GroupNodeInfo", string(data)}, " ")
}
