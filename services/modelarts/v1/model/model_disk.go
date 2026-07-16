package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Disk 磁盘信息。
type Disk struct {

	// 磁盘大小。
	Size *string `json:"size,omitempty"`

	// 磁盘大小单位，一般为GB。
	Unit *string `json:"unit,omitempty"`
}

func (o Disk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Disk struct{}"
	}

	return strings.Join([]string{"Disk", string(data)}, " ")
}
