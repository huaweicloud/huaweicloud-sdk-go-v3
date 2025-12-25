package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Directory directory 目录分组
type Directory struct {
}

func (o Directory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Directory struct{}"
	}

	return strings.Join([]string{"Directory", string(data)}, " ")
}
