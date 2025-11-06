package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectoryDto struct {

	// 目录ID。
	DirectoryId *string `json:"directory_id,omitempty"`

	// 目录名称。
	DirectoryName *string `json:"directory_name,omitempty"`
}

func (o DirectoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectoryDto struct{}"
	}

	return strings.Join([]string{"DirectoryDto", string(data)}, " ")
}
