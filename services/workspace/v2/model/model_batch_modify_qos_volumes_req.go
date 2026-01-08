package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyQosVolumesReq 批量修改磁盘QOS请求。
type BatchModifyQosVolumesReq struct {

	// 修改QOS磁盘ids。
	VolumeIds []string `json:"volume_ids"`

	Qos *Qos `json:"qos"`
}

func (o BatchModifyQosVolumesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyQosVolumesReq struct{}"
	}

	return strings.Join([]string{"BatchModifyQosVolumesReq", string(data)}, " ")
}
