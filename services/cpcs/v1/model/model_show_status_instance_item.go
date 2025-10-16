package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatusInstanceItem struct {

	// 实例列表
	InstanceList *[]ShowStatusInstanceItemInstanceList `json:"instance_list,omitempty"`

	// 采集时间，UNIX时间戳，单位毫秒
	Timestamp *int32 `json:"timestamp,omitempty"`
}

func (o ShowStatusInstanceItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusInstanceItem struct{}"
	}

	return strings.Join([]string{"ShowStatusInstanceItem", string(data)}, " ")
}
