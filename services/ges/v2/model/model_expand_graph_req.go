package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandGraphReq 扩副本请求体
type ExpandGraphReq struct {
	Expand *ExpandGraphReqExpand `json:"expand"`
}

func (o ExpandGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphReq struct{}"
	}

	return strings.Join([]string{"ExpandGraphReq", string(data)}, " ")
}
