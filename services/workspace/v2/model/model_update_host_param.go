package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateHostParam struct {

	// 自动放置，off-取消自动放置，on-打开自动放置。
	AutoPlacement *string `json:"auto_placement,omitempty"`

	// 云办公主机名称。
	Name *string `json:"name,omitempty"`

	// 云办公主机id。
	HostId *string `json:"host_id,omitempty"`
}

func (o UpdateHostParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostParam struct{}"
	}

	return strings.Join([]string{"UpdateHostParam", string(data)}, " ")
}
