package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestBackUpMindmapParam struct {
	Params *BackUpMindmapParam `json:"params"`
}

func (o CommRequestBackUpMindmapParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestBackUpMindmapParam struct{}"
	}

	return strings.Join([]string{"CommRequestBackUpMindmapParam", string(data)}, " ")
}
