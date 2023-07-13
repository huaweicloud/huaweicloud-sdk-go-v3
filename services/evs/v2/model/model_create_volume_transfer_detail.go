package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeTransferDetail struct {

	// 云硬盘过户的身份认证密钥。
	AuthKey string `json:"auth_key"`

	// 云硬盘过户记录的创建时间。  时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	CreatedAt string `json:"created_at"`

	// 云硬盘过户记录的ID。
	Id string `json:"id"`

	// 云硬盘过户记录的链接。
	Links []Link `json:"links"`

	// 云硬盘过户记录的名称。
	Name string `json:"name"`

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`
}

func (o CreateVolumeTransferDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeTransferDetail struct{}"
	}

	return strings.Join([]string{"CreateVolumeTransferDetail", string(data)}, " ")
}
