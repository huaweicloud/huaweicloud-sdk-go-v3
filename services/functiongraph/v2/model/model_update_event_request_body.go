package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEventRequestBody struct {

	// 测试事件content，为json字符串的base64编码
	Content string `json:"content"`
}

func (o UpdateEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEventRequestBody", string(data)}, " ")
}
