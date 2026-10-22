package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomAnswersResponse Response Object
type ShowCustomAnswersResponse struct {

	// 代答库ID
	Id *string `json:"id,omitempty"`

	// 代答库名称
	Name *string `json:"name,omitempty"`

	// 代答库描述
	Description *string `json:"description,omitempty"`

	// 答案列表（按seq排序）
	Answers *[]string `json:"answers,omitempty"`

	// 引用的策略列表
	ReferencedRules *[]ReferencedRule `json:"referenced_rules,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *int64 `json:"update_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowCustomAnswersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomAnswersResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomAnswersResponse", string(data)}, " ")
}
