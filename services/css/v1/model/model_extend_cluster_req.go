package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendClusterReq struct {
	Grow *ExtendClusterGrowReq `json:"grow"`

	IsAutoPay *int32 `json:"isAutoPay,omitempty"`
}

func (o ExtendClusterReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendClusterReq struct{}"
	}

	return strings.Join([]string{"ExtendClusterReq", string(data)}, " ")
}
