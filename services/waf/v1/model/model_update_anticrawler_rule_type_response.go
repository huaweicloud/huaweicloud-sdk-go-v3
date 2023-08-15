package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAnticrawlerRuleTypeResponse Response Object
type UpdateAnticrawlerRuleTypeResponse struct {

	// JS脚本反爬虫规则类型，指定防护路径：anticrawler_specific_url 排除防护路径：anticrawler_except_url
	AnticrawlerType *string `json:"anticrawler_type,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateAnticrawlerRuleTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnticrawlerRuleTypeResponse struct{}"
	}

	return strings.Join([]string{"UpdateAnticrawlerRuleTypeResponse", string(data)}, " ")
}
