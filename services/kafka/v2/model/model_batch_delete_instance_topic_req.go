package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteInstanceTopicReq struct {
	// 待删除的topic列表。

	Topics *[]string `json:"topics,omitempty"`
}

func (o BatchDeleteInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteInstanceTopicReq", string(data)}, " ")
}
