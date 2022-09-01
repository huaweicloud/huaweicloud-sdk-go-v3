package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpandGraphReq struct {
	Expand *ReplicationReq `json:"expand" xml:"expand"`
}

func (o ExpandGraphReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphReq struct{}"
	}

	return strings.Join([]string{"ExpandGraphReq", string(data)}, " ")
}
