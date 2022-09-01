package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMqsInstanceTopicReq struct {

	// 待删除的topic列表。
	Topics []string `json:"topics" xml:"topics"`
}

func (o BatchDeleteMqsInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMqsInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteMqsInstanceTopicReq", string(data)}, " ")
}
