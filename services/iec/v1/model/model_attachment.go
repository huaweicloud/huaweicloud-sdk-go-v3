package model

import (
	"encoding/json"

	"strings"
)

//
type Attachment struct {
	// 挂载信息对应的ID。

	AttachmentId string `json:"attachment_id"`
	// 挂载点。

	Device string `json:"device"`
	// 边缘硬盘挂载到的边缘实例对应的物理主机的名称。

	HostName string `json:"host_name"`
	// 挂载的资源ID。

	Id string `json:"id"`
	// 硬盘挂载到的边缘实例的ID。

	ServerId string `json:"server_id"`
	// 磁盘ID。

	VolumeId string `json:"volume_id"`
}

func (o Attachment) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Attachment struct{}"
	}

	return strings.Join([]string{"Attachment", string(data)}, " ")
}
