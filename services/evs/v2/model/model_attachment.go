package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云硬盘挂载信息。
type Attachment struct {
	// 挂载的时间信息。  时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX

	AttachedAt string `json:"attached_at"`
	// 挂载信息对应的ID。

	AttachmentId string `json:"attachment_id"`
	// 挂载点。

	Device string `json:"device"`
	// 云硬盘挂载到的云服务器对应的物理主机的名称。

	HostName string `json:"host_name"`
	// 挂载的资源ID。

	Id string `json:"id"`
	// 云硬盘挂载到的云服务器的 ID。

	ServerId string `json:"server_id"`
	// 云硬盘ID。

	VolumeId string `json:"volume_id"`
}

func (o Attachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attachment struct{}"
	}

	return strings.Join([]string{"Attachment", string(data)}, " ")
}
