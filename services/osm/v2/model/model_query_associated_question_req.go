package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryAssociatedQuestionReq struct {

	// 用户输入问题
	Question string `json:"question"`

	// 返回匹配度最高的数据条数
	Top *int32 `json:"top,omitempty"`

	// 问答领域
	Domains *[]string `json:"domains,omitempty"`
}

func (o QueryAssociatedQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryAssociatedQuestionReq struct{}"
	}

	return strings.Join([]string{"QueryAssociatedQuestionReq", string(data)}, " ")
}
