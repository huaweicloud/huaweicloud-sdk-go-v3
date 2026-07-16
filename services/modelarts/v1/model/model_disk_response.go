package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskResponse 磁盘信息。
type DiskResponse struct {

	// 磁盘大小。
	Size *int32 `json:"size,omitempty"`

	// 磁盘大小单位。
	Unit *string `json:"unit,omitempty"`
}

func (o DiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskResponse struct{}"
	}

	return strings.Join([]string{"DiskResponse", string(data)}, " ")
}
