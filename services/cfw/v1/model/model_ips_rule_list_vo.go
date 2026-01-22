package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsRuleListVo struct {

	// 防火墙ID
	FwInstanceId *string `json:"fw_instance_id,omitempty"`

	// 每页显示个数，范围为1-1024
	Limit *int32 `json:"limit,omitempty"`

	// 防护对象id
	ObjectId *string `json:"object_id,omitempty"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 查询ips规则列表
	Records *[]IpsRuleVo `json:"records,omitempty"`

	// 查询ips规则总数
	Total *int32 `json:"total,omitempty"`
}

func (o IpsRuleListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleListVo struct{}"
	}

	return strings.Join([]string{"IpsRuleListVo", string(data)}, " ")
}
