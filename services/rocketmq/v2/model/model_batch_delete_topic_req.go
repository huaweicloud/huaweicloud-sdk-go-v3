package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteTopicReq struct {

	// 主题列表，当批量删除主题时使用。
	Topics *[]string `json:"topics,omitempty" xml:"topics"`
}

func (o BatchDeleteTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteTopicReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteTopicReq", string(data)}, " ")
}
