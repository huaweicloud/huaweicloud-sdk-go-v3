package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAnticrawlerRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 匹配条件列表
	Conditions *[]AnticrawlerCondition `json:"conditions,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// JS脚本反爬虫规则类型，指定防护路径：anticrawler_specific_url 排除防护路径：anticrawler_except_url
	Type *string `json:"type,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则状态，0：关闭，1：开启
	Status *int32 `json:"status,omitempty"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到1000。
	Priority       *int32 `json:"priority,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeleteAnticrawlerRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnticrawlerRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnticrawlerRuleResponse", string(data)}, " ")
}
