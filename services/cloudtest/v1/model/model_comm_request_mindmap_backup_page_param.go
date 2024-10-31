package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommRequestMindmapBackupPageParam struct {
	Params *MindmapBackupPageParam `json:"params"`
}

func (o CommRequestMindmapBackupPageParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommRequestMindmapBackupPageParam struct{}"
	}

	return strings.Join([]string{"CommRequestMindmapBackupPageParam", string(data)}, " ")
}
