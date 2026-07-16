package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobVolumeResp 训练作业挂载卷信息。
type JobVolumeResp struct {
	Nfs *NfsResp `json:"nfs,omitempty"`
}

func (o JobVolumeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobVolumeResp struct{}"
	}

	return strings.Join([]string{"JobVolumeResp", string(data)}, " ")
}
