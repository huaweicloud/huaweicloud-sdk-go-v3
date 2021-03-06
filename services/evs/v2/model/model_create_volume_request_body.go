package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateVolumeRequestBody struct {
	BssParam *BssParamForCreateVolume `json:"bssParam,omitempty"`

	Volume *CreateVolumeOption `json:"volume"`
	// 创建云硬盘并挂载到目标虚拟机。

	ServerId *string `json:"server_id,omitempty"`

	OSSCHHNTschedulerHints *CreateVolumeSchedulerHints `json:"OS-SCH-HNT:scheduler_hints,omitempty"`
}

func (o CreateVolumeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateVolumeRequestBody struct{}"
	}

	return strings.Join([]string{"CreateVolumeRequestBody", string(data)}, " ")
}
