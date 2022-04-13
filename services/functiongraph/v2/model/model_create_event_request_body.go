package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建测试事件请求体。
type CreateEventRequestBody struct {
	// 测试事件名称。只能由字母、数字、中划线和下划线组成，且必须以大写或小写字母开头。

	Name string `json:"name"`
	// 测试事件content。

	Content string `json:"content"`
}

func (o CreateEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEventRequestBody", string(data)}, " ")
}
