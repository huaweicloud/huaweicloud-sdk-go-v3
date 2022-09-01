package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改源端信息json的请求体，当前只支持修改源端服务器名称和迁移项目id
type PutSourceServerBody struct {

	// 源端服务器修改后的名字
	Name *string `json:"name,omitempty" xml:"name"`

	// 源端服务器修改后所属的迁移项目id
	Migprojectid *string `json:"migprojectid,omitempty" xml:"migprojectid"`

	// 磁盘
	Disks *[]PutDisk `json:"disks,omitempty" xml:"disks"`

	// 卷组
	VolumeGroups *[]PutVolumeGroups `json:"volume_groups,omitempty" xml:"volume_groups"`
}

func (o PutSourceServerBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutSourceServerBody struct{}"
	}

	return strings.Join([]string{"PutSourceServerBody", string(data)}, " ")
}
