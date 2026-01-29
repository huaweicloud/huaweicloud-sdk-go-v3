package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutSourceServerBody 修改源端信息json的请求体
type PutSourceServerBody struct {

	// 源端服务器修改后的名字
	Name *string `json:"name,omitempty"`

	// 源端服务器修改后所属的迁移项目ID
	Migprojectid *string `json:"migprojectid,omitempty"`

	// 磁盘，仅在“待配置目的端”状态下，此修改才生效
	Disks *[]PutDisk `json:"disks,omitempty"`

	// 卷组，仅在“待配置目的端”状态下，此修改才生效
	VolumeGroups *[]PutVolumeGroups `json:"volume_groups,omitempty"`
}

func (o PutSourceServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutSourceServerBody struct{}"
	}

	return strings.Join([]string{"PutSourceServerBody", string(data)}, " ")
}
