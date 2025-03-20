package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServerVolumeAttachment struct {

	// 云硬盘挂载盘符，即磁盘挂载点。
	Device *string `json:"device,omitempty"`

	// 挂载ID，与云硬盘ID相同，UUID格式。
	Id *string `json:"id,omitempty"`

	// 弹性云服务器ID，UUID格式。
	ServerId *string `json:"serverId,omitempty"`

	// 云硬盘ID，UUID格式。
	VolumeId *string `json:"volumeId,omitempty"`
}

func (o ServerVolumeAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerVolumeAttachment struct{}"
	}

	return strings.Join([]string{"ServerVolumeAttachment", string(data)}, " ")
}
