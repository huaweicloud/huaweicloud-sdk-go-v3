package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 运行日志单个条目
type RunlogItem struct {

	// 日志的唯一标识
	Id *string `json:"id,omitempty" xml:"id"`

	// 运行日志文件名
	FileName *string `json:"file_name,omitempty" xml:"file_name"`

	// 分片名称
	GroupName *string `json:"group_name,omitempty" xml:"group_name"`

	// 采集运行日志所在副本的IP
	ReplicationIp *string `json:"replication_ip,omitempty" xml:"replication_ip"`

	// 获取运行日志状态
	Status *string `json:"status,omitempty" xml:"status"`

	// 运行日志采集的日期，格式为\"yyyy-MM-dd\"
	Time *string `json:"time,omitempty" xml:"time"`
}

func (o RunlogItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunlogItem struct{}"
	}

	return strings.Join([]string{"RunlogItem", string(data)}, " ")
}
