package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainQuotaResponseQuotas struct {

	// **参数解释：** 资源类型。 **取值范围：** - zone：公网域名配额 - private_zone：内网域名配额 - record_set：记录集配额 - url_record：显性/隐性URL记录集配额 - ptr_record：反向解析配额 - custom_line：自定义线路配额 - line_group：线路分组配额 - inbound_endpoint：入站终端节点配额 - outbound_endpoint：出站终端节点配额 - resolver_rule：转发规则配额
	QuotaKey string `json:"quota_key"`

	// 资源配额的最大值。
	QuotaLimit int32 `json:"quota_limit"`

	// 配额已使用数量。
	Used int32 `json:"used"`

	// 配额统计单位，取固定值“count”。
	Unit string `json:"unit"`
}

func (o DomainQuotaResponseQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainQuotaResponseQuotas struct{}"
	}

	return strings.Join([]string{"DomainQuotaResponseQuotas", string(data)}, " ")
}
