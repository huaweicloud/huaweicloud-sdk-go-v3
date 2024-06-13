package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestMindmapPageParamV3 struct {
	Params *MindmapPageParamV3 `json:"params,omitempty"`
}

func (o CommRequestMindmapPageParamV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestMindmapPageParamV3 struct{}"
	}

	return strings.Join([]string{"CommRequestMindmapPageParamV3", string(data)}, " ")
}
