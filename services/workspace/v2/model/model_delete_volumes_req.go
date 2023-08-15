package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVolumesReq 删除桌面数据盘请求。
type DeleteVolumesReq struct {

	// 待删除的桌面数据盘ID列表。
	VolumeIds *[]string `json:"volume_ids,omitempty"`
}

func (o DeleteVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVolumesReq struct{}"
	}

	return strings.Join([]string{"DeleteVolumesReq", string(data)}, " ")
}
