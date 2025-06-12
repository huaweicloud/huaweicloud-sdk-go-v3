package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateRecycleBinPolicyOption struct {

	// 虚拟机进入回收站后多久删除
	RetentionHour *int32 `json:"retention_hour,omitempty"`

	// 创建多久的虚拟机可以进入回收站
	RecycleThresholdDay *int32 `json:"recycle_threshold_day,omitempty"`
}

func (o UpdateRecycleBinPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecycleBinPolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateRecycleBinPolicyOption", string(data)}, " ")
}
