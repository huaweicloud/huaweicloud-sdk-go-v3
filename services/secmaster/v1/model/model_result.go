package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Result struct {

	// 平均消费大小
	AverageMsgBytes float32 `json:"average_msg_bytes"`

	// 消费条数
	SubscribeMsgs int64 `json:"subscribe_msgs"`
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
