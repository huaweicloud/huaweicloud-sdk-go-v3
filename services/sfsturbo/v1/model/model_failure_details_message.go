package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FailureDetailsMessage 数据存储库生命周期中出现的错误信息
type FailureDetailsMessage struct {

	// 错误信息
	Message string `json:"message"`
}

func (o FailureDetailsMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailureDetailsMessage struct{}"
	}

	return strings.Join([]string{"FailureDetailsMessage", string(data)}, " ")
}
