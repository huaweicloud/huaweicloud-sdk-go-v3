package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IpsRuleChangeRespBody struct {

	// 错误代码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 分组id
	GroupId *string `json:"group_id,omitempty"`

	// 防火墙ID
	Id *string `json:"id,omitempty"`

	// ips的id列表
	IpsIds *[]string `json:"ips_ids,omitempty"`

	// 修改结果
	Result *bool `json:"result,omitempty"`
}

func (o IpsRuleChangeRespBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsRuleChangeRespBody struct{}"
	}

	return strings.Join([]string{"IpsRuleChangeRespBody", string(data)}, " ")
}
