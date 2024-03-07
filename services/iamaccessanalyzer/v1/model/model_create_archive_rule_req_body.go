package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateArchiveRuleReqBody struct {

	// 匹配要返回的访问分析结果的筛选器。
	Filters []FindingFilter `json:"filters"`

	// 创建存档规则的名称。
	Name string `json:"name"`
}

func (o CreateArchiveRuleReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArchiveRuleReqBody struct{}"
	}

	return strings.Join([]string{"CreateArchiveRuleReqBody", string(data)}, " ")
}
