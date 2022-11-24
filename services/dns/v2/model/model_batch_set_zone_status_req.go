package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSetZoneStatusReq struct {

	// 待设置Zone状态，当前仅支持DISABLE或ENABLE。
	Status *string `json:"status,omitempty"`

	// 待设置Zone状态，当前仅支持DISABLE或ENABLE。
	ZoneIds *[]string `json:"zone_ids,omitempty"`
}

func (o BatchSetZoneStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetZoneStatusReq struct{}"
	}

	return strings.Join([]string{"BatchSetZoneStatusReq", string(data)}, " ")
}
