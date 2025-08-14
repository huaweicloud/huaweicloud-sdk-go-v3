package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBaselineWhiteListRequestBody 新增基线白名单请求体信息
type AddBaselineWhiteListRequestBody struct {

	// 基线检查的操作系统 - Linux - Windows
	OsType string `json:"os_type"`

	// 基线检查白名单的规则范围 - specific_host：部分主机 - all_host：全部主机
	RuleType string `json:"rule_type"`

	// rule_type为specific_host时，该字段为待添加的白名单主机id列表，rule_type为all_host时无该字段
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 待添加的白名单的检查项列表
	RuleList []AddBaselineWhiteListRequestBodyRuleList `json:"rule_list"`

	// 基线白名单备注
	Description *string `json:"description,omitempty"`
}

func (o AddBaselineWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBaselineWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"AddBaselineWhiteListRequestBody", string(data)}, " ")
}
