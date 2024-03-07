package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateArchiveRuleReqBody struct {

	// 匹配要返回的访问分析结果的筛选器。
	Filters []FindingFilter `json:"filters"`
}

func (o UpdateArchiveRuleReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArchiveRuleReqBody struct{}"
	}

	return strings.Join([]string{"UpdateArchiveRuleReqBody", string(data)}, " ")
}
