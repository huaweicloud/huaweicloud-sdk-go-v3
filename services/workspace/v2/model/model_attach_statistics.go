package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachStatistics 分配情况统计
type AttachStatistics struct {

	// 已分配个数。
	AttachedNum *int32 `json:"attached_num,omitempty"`

	// 未分配个数。
	UnattachedNum *int32 `json:"unattached_num,omitempty"`

	// 分配中个数。
	AttachingNum *int32 `json:"attaching_num,omitempty"`

	// 分配失败的个数。
	AttachErrorNum *int32 `json:"attach_error_num,omitempty"`
}

func (o AttachStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachStatistics struct{}"
	}

	return strings.Join([]string{"AttachStatistics", string(data)}, " ")
}
