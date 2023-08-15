package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionProgress struct {

	// 创建文件系统的进度。
	Creating *string `json:"CREATING,omitempty"`
}

func (o ActionProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionProgress struct{}"
	}

	return strings.Join([]string{"ActionProgress", string(data)}, " ")
}
