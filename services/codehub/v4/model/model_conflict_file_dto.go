package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConflictFileDto struct {

	// 旧文件
	OldPath *string `json:"old_path,omitempty"`

	// 新文件
	NewPath *string `json:"new_path,omitempty"`
}

func (o ConflictFileDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConflictFileDto struct{}"
	}

	return strings.Join([]string{"ConflictFileDto", string(data)}, " ")
}
