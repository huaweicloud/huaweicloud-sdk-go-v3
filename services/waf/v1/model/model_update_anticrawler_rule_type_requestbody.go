package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAnticrawlerRuleTypeRequestbody struct {

	// JS脚本反爬虫规则类型，指定防护路径：anticrawler_specific_url 排除防护路径：anticrawler_except_url
	AnticrawlerType string `json:"anticrawler_type"`
}

func (o UpdateAnticrawlerRuleTypeRequestbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleTypeRequestbody struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleTypeRequestbody", string(data)}, " ")
}
