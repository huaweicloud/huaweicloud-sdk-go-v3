package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBody struct {

	// 需要删除的资源名称列表
	Name []string `json:"name"`
}

func (o BatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBody", string(data)}, " ")
}
