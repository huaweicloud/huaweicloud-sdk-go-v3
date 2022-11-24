package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 保护实例挂载复制对请求数据结构
type ProtectedInstanceAttachReplicationRequestParams struct {

	// 复制对的ID。
	ReplicationId string `json:"replication_id"`

	// 复制对中使用的磁盘挂载点。  说明:新增加的磁盘挂载点不能和已有的磁盘挂载点相同。 对于采用XEN虚拟化类型的弹性云服务器组成的保护实例，系统盘挂载点请指定/dev/sda；数据盘挂载点请按英文字母顺序依次指定，如/dev/sdb，/dev/sdc，如果指定了以“/dev/vd”开头的挂载点，系统默认改为“/dev/sd”。对于采用KVM虚拟化类型的弹性云服务器组成的保护实例，系统盘挂载点请指定/dev/vda；数据盘挂载点请按英文字母顺序依次指定，如/dev/vdb，/dev/vdc，如果指定了以“/dev/sd”开头的挂载点，系统默认改为“/dev/vd”。
	Device string `json:"device"`
}

func (o ProtectedInstanceAttachReplicationRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectedInstanceAttachReplicationRequestParams struct{}"
	}

	return strings.Join([]string{"ProtectedInstanceAttachReplicationRequestParams", string(data)}, " ")
}
