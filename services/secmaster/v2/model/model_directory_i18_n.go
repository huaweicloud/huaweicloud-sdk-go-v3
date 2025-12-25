package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DirectoryI18N struct {

	// directory 目录分组
	Directory *string `json:"directory,omitempty"`

	// directory 目录分组
	DirectoryEn *string `json:"directory_en,omitempty"`

	// directory 目录分组
	DirectoryFr *string `json:"directory_fr,omitempty"`
}

func (o DirectoryI18N) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectoryI18N struct{}"
	}

	return strings.Join([]string{"DirectoryI18N", string(data)}, " ")
}
