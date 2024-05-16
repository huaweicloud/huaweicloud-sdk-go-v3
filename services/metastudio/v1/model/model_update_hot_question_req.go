package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotQuestionReq 修改热点问题请求。
type UpdateHotQuestionReq struct {

	// 热点问题。
	HotQuestion *string `json:"hot_question,omitempty"`
}

func (o UpdateHotQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotQuestionReq struct{}"
	}

	return strings.Join([]string{"UpdateHotQuestionReq", string(data)}, " ")
}
