package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建文件系统参数body
type CreateShareRequestBody struct {
	Share *Share `json:"share"`
}

func (o CreateShareRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareRequestBody struct{}"
	}

	return strings.Join([]string{"CreateShareRequestBody", string(data)}, " ")
}
