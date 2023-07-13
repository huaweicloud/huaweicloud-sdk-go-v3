package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebuildImageRequestBody 重试镜像任务请求体。
type RebuildImageRequestBody struct {

	// 重试镜像任务动作，取值retry。
	Action string `json:"action"`
}

func (o RebuildImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebuildImageRequestBody struct{}"
	}

	return strings.Join([]string{"RebuildImageRequestBody", string(data)}, " ")
}
