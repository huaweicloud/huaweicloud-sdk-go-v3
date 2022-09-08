package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeTransfer struct {

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

func (o VolumeTransfer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTransfer struct{}"
	}

	return strings.Join([]string{"VolumeTransfer", string(data)}, " ")
}
