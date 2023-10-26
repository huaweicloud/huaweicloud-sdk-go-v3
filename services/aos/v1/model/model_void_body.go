package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VoidBody 空响应体
type VoidBody struct {
}

func (o VoidBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoidBody struct{}"
	}

	return strings.Join([]string{"VoidBody", string(data)}, " ")
}
