package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SessionAskQuestionReq struct {

	// 用户输入问题
	Question string `json:"question"`

	// 问答领域
	Domain *string `json:"domain,omitempty"`

	// 最大返回数据条数
	Top *int32 `json:"top,omitempty"`

	// 主题列表
	Themes *[]RelationTheme `json:"themes,omitempty"`

	// 语料ID
	SourceQaPairId *string `json:"source_qa_pair_id,omitempty"`

	// 是否需要兜底
	ChatEnable *bool `json:"chat_enable,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`
}

func (o SessionAskQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SessionAskQuestionReq struct{}"
	}

	return strings.Join([]string{"SessionAskQuestionReq", string(data)}, " ")
}
