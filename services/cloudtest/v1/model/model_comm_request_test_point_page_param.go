package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestTestPointPageParam struct {
	Params *TestPointPageParam `json:"params,omitempty"`
}

func (o CommRequestTestPointPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestTestPointPageParam struct{}"
	}

	return strings.Join([]string{"CommRequestTestPointPageParam", string(data)}, " ")
}
