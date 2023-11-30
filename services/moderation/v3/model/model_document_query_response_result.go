package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DocumentQueryResponseResult 作业审核结果，当作业状态为succeeded时存在。
type DocumentQueryResponseResult struct {

	// 文档审核结果是否通过。 block：包含敏感信息，不通过 review：需要人工复检 pass：不包含敏感信息，通过
	Suggestion *string `json:"suggestion,omitempty"`

	// 审核详情
	Details *[]DocumentQueryResponseResultDetails `json:"details,omitempty"`
}

func (o DocumentQueryResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DocumentQueryResponseResult struct{}"
	}

	return strings.Join([]string{"DocumentQueryResponseResult", string(data)}, " ")
}
