package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChatResultRequestMessage 获取异步推断结果的请求
type ChatResultRequestMessage struct {

	// request id
	ResultId string `json:"result_id"`
}

func (o ChatResultRequestMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatResultRequestMessage struct{}"
	}

	return strings.Join([]string{"ChatResultRequestMessage", string(data)}, " ")
}
