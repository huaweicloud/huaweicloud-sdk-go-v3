package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandGraphReqExpand expand是一个对象
type ExpandGraphReqExpand struct {

	// 新扩副本数量。
	Replication int32 `json:"replication"`
}

func (o ExpandGraphReqExpand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphReqExpand struct{}"
	}

	return strings.Join([]string{"ExpandGraphReqExpand", string(data)}, " ")
}
