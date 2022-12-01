package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 磁盘
type Volume struct {

	// 磁盘名称，取值范围为 SSD（超高IO）,高IO（SAS），普通IO（SATA）
	Volume *string `json:"volume,omitempty"`

	// 磁盘容量
	Capacity *int32 `json:"capacity,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
