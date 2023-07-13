package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAnticrawlerRuleRequestBody struct {

	// 匹配条件列表
	Conditions []AnticrawlerCondition `json:"conditions"`

	// 规则名称
	Name string `json:"name"`

	// JS脚本反爬虫规则类型，指定防护路径：anticrawler_specific_url 排除防护路径：anticrawler_except_url
	Type string `json:"type"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到1000。
	Priority int32 `json:"priority"`
}

func (o UpdateAnticrawlerRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleRequestBody", string(data)}, " ")
}
