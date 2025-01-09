package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddVolumes 磁盘订单请求。
type AddVolumes struct {

	// 服务ID。
	ServiceId *string `json:"service_id,omitempty"`

	// 订单的磁盘信息列表。
	Volumes *[]Volume `json:"volumes,omitempty"`
}

func (o AddVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVolumes struct{}"
	}

	return strings.Join([]string{"AddVolumes", string(data)}, " ")
}
