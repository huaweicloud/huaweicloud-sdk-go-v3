package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateResponseBody struct {
	// 创建出的服务作业ID

	Id string `json:"id"`
}

func (o CreateResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResponseBody struct{}"
	}

	return strings.Join([]string{"CreateResponseBody", string(data)}, " ")
}
