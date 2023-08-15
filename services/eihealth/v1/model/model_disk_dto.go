package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiskDto struct {

	// 磁盘类型
	Type *string `json:"type,omitempty"`

	// 磁盘大小
	Space *string `json:"space,omitempty"`
}

func (o DiskDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskDto struct{}"
	}

	return strings.Join([]string{"DiskDto", string(data)}, " ")
}
