package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandGraphReq struct {
	Expand *ReplicationReq `json:"expand"`
}

func (o ExpandGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphReq struct{}"
	}

	return strings.Join([]string{"ExpandGraphReq", string(data)}, " ")
}
