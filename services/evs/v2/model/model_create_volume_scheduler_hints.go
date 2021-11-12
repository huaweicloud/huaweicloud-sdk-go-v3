package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建云硬盘的调度参数。
type CreateVolumeSchedulerHints struct {
	// 指定专属存储池ID，表示将云硬盘创建在该ID对应的存储池中。

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`
}

func (o CreateVolumeSchedulerHints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeSchedulerHints struct{}"
	}

	return strings.Join([]string{"CreateVolumeSchedulerHints", string(data)}, " ")
}
