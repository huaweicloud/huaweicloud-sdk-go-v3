package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CheckpointExtraInfoResp struct {

	// 备份名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 备份描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 备份保留天数
	RetentionDuration *int32 `json:"retention_duration,omitempty" xml:"retention_duration"`
}

func (o CheckpointExtraInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckpointExtraInfoResp struct{}"
	}

	return strings.Join([]string{"CheckpointExtraInfoResp", string(data)}, " ")
}
