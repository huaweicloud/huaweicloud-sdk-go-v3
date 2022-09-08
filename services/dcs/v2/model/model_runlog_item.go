package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 运行日志单个条目
type RunlogItem struct {

	// 日志的唯一标识
	Id *string `json:"id,omitempty"`

	// 运行日志文件名
	FileName *string `json:"file_name,omitempty"`

	// 分片名称
	GroupName *string `json:"group_name,omitempty"`

	// 采集运行日志所在副本的IP
	ReplicationIp *string `json:"replication_ip,omitempty"`

	// 获取运行日志状态
	Status *string `json:"status,omitempty"`

	// 运行日志采集的日期，格式为\"yyyy-MM-dd\"
	Time *string `json:"time,omitempty"`

	// 日志文件的ID
	BackupId *string `json:"backupId,omitempty"`
}

func (o RunlogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunlogItem struct{}"
	}

	return strings.Join([]string{"RunlogItem", string(data)}, " ")
}
