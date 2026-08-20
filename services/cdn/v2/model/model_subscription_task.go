package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionTask 订阅任务
type SubscriptionTask struct {

	// **参数解释：** 订阅任务id **取值范围：** 不涉及
	Id *int64 `json:"id,omitempty"`

	// **参数解释：** 创建时间 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 最近更新时间 **取值范围：** 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释：** 订阅任务的名称 **约束限制：** 长度不超过32个字符。 **取值范围：** 支持大小写字母、中文字符、“-”
	Name *string `json:"name,omitempty"`

	// **参数解释：** 订阅任务类型 **取值范围：** - 0：日报 - 1：周报 - 2：月报
	PeriodType *int32 `json:"period_type,omitempty"`

	// **参数解释：** 接收运营报表的邮箱地址 > 支持同时输入多个邮箱地址  **取值范围：** 多个邮箱地址用英文逗号（,）分隔
	Emails *string `json:"emails,omitempty"`

	// **参数解释：** 订阅的域名列表 > 支持同时输入多个域名  **取值范围：** - 多个域名用半角逗号（,）分隔 - 如果该参数为all，则为账号下的所有域名订阅运营报表
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 运营报表类型 > 支持同时输入多个报表类型  **取值范围：** - 多个报表类型以英文逗号（,）分隔 - 0：访问区域分布 - 1：国家分布 - 2：运营商分布 - 3：域名排行（按流量排序） - 4：热门URL（按流量排序） - 5：热门URL（按请求数排序） - 6：热门Referer（按流量排序） - 7：热门Referer（按请求数排序） - 10：回源热门URL（按流量排序） - 11：回源热门URL（按请求数排序） - 13：热门UA（按流量排序） - 14：热门UA（按请求数排序）
	ReportType *string `json:"report_type,omitempty"`
}

func (o SubscriptionTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTask struct{}"
	}

	return strings.Join([]string{"SubscriptionTask", string(data)}, " ")
}
