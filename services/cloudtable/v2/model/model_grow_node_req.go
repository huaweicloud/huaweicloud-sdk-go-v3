package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GrowNodeReq struct {

	// 扩容节点类型：rs,tsdb,lemon
	ComponentName string `json:"component_name"`

	// 扩容节点范围是 [2,10]
	NodeNum int32 `json:"node_num"`
}

func (o GrowNodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrowNodeReq struct{}"
	}

	return strings.Join([]string{"GrowNodeReq", string(data)}, " ")
}
