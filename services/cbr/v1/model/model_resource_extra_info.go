package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceExtraInfo struct {
	// 需要排除备份的卷id。仅虚拟机磁盘级备份有效

	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty"`
	// 要备份的卷

	IncludeVolumes *[]ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty"`
}

func (o ResourceExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfo struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfo", string(data)}, " ")
}
