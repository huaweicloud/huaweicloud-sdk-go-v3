package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VolumeTransferSummary struct {

	// 云硬盘过户记录的ID。
	Id string `json:"id"`

	// 云硬盘过户记录的链接
	Links []Link `json:"links"`

	// 云硬盘过户记录的名称
	Name string `json:"name"`

	// 云硬盘ID。
	VolumeId string `json:"volume_id"`
}

func (o VolumeTransferSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTransferSummary struct{}"
	}

	return strings.Join([]string{"VolumeTransferSummary", string(data)}, " ")
}
