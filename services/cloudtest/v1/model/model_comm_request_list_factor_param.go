package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestListFactorParam struct {
	Params *ListFactorParam `json:"params"`
}

func (o CommRequestListFactorParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestListFactorParam struct{}"
	}

	return strings.Join([]string{"CommRequestListFactorParam", string(data)}, " ")
}
