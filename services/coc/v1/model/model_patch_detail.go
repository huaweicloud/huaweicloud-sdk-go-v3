package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PatchDetail struct {

	// 安装时间
	InstalledTime *int64 `json:"installed_time,omitempty"`

	// 补丁基线id
	PatchBaselineId *string `json:"patch_baseline_id,omitempty"`

	// 补丁基线名称
	PatchBaselineName *string `json:"patch_baseline_name,omitempty"`

	// 补丁状态
	PatchStatus *string `json:"patch_status,omitempty"`
}

func (o PatchDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchDetail struct{}"
	}

	return strings.Join([]string{"PatchDetail", string(data)}, " ")
}
