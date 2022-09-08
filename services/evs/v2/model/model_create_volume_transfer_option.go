package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVolumeTransferOption struct {

	// 云硬盘过户记录的名称。最大支持255个字节。
	Name string `json:"name"`

	// 云硬盘ID。  通过[查询所有云硬盘详情](https://support.huaweicloud.com/api-evs/evs_04_3033.html)获取。
	VolumeId string `json:"volume_id"`
}

func (o CreateVolumeTransferOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeTransferOption struct{}"
	}

	return strings.Join([]string{"CreateVolumeTransferOption", string(data)}, " ")
}
