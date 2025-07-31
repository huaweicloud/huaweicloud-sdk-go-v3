package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBaselineWhiteListRequestBody 修改基线白名单请求体信息
type ChangeBaselineWhiteListRequestBody struct {

	// 被修改的白名单id
	Id string `json:"id"`

	// 基线检查白名单的规则范围 - specific_host 部分主机 - all_host      全部主机
	RuleType string `json:"rule_type"`

	// rule_type为specific_host时,该字段为待修改的白名单主机id列表, rule_type为all_host时无该字段
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// 基线白名单备注
	Description *string `json:"description,omitempty"`
}

func (o ChangeBaselineWhiteListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaselineWhiteListRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeBaselineWhiteListRequestBody", string(data)}, " ")
}
