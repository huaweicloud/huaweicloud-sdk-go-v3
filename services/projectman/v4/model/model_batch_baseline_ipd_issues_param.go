package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchBaselineIpdIssuesParam struct {

	// 需要基线的工作项ID数组。可以通过查询工作项列表或者查询树状工作项接口获取，响应消息体中的id字段的值就是工作项ID。
	Id []string `json:"id"`

	Attribute *BatchBaselineIpdIssuesParamAttribute `json:"attribute"`
}

func (o BatchBaselineIpdIssuesParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBaselineIpdIssuesParam struct{}"
	}

	return strings.Join([]string{"BatchBaselineIpdIssuesParam", string(data)}, " ")
}
