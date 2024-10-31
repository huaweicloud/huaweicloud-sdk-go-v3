package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestMindmapRecyclePageParam struct {
	Params *MindmapRecyclePageParam `json:"params"`
}

func (o CommRequestMindmapRecyclePageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestMindmapRecyclePageParam struct{}"
	}

	return strings.Join([]string{"CommRequestMindmapRecyclePageParam", string(data)}, " ")
}
