package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetJobInfoResponseBodyJobEntitiesVolume 调整实例容量任务时返回。
type GetJobInfoResponseBodyJobEntitiesVolume struct {

	// 磁盘类型。
	Type *string `json:"type,omitempty"`

	// 实例原本的磁盘大小（单位：GB）。
	OriginalSize *string `json:"original_size,omitempty"`

	// 任务的目标磁盘大小（单位：GB）。
	TargetSize *string `json:"target_size,omitempty"`
}

func (o GetJobInfoResponseBodyJobEntitiesVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetJobInfoResponseBodyJobEntitiesVolume struct{}"
	}

	return strings.Join([]string{"GetJobInfoResponseBodyJobEntitiesVolume", string(data)}, " ")
}
