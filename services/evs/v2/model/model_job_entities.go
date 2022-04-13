package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Job的响应信息。
type JobEntities struct {
	// 云硬盘的类型。

	VolumeType *string `json:"volume_type,omitempty"`
	// 云硬盘的容量，单位为GB。

	Size *int32 `json:"size,omitempty"`
	// 云硬盘的ID。

	VolumeId *string `json:"volume_id,omitempty"`
	// 云硬盘的名称。

	Name *string `json:"name,omitempty"`
	// 子Job的信息。当存在子Job信息时，entities中的其他字段将不会返回。

	SubJobs *[]SubJob `json:"sub_jobs,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
