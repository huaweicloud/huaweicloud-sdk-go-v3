package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QaGraphAnswer struct {

	// 答案
	Answer *string `json:"answer,omitempty"`

	// 评分
	Score *float64 `json:"score,omitempty"`

	// 类型
	Type *int32 `json:"type,omitempty"`

	// 列表
	Options *[]string `json:"options,omitempty"`
}

func (o QaGraphAnswer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaGraphAnswer struct{}"
	}

	return strings.Join([]string{"QaGraphAnswer", string(data)}, " ")
}
