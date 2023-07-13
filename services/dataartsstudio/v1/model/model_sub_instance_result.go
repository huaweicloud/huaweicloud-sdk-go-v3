package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubInstanceResult struct {

	// 子规则ID
	SubRuleId *string `json:"sub_rule_id,omitempty"`

	// 子规则实例ID
	SubInstanceInstanceId *string `json:"sub_instance_instance_id,omitempty"`

	// 异常表任务状态 UNSUPPORTED:不支持异常表,READY：准备中,RUNNING：运行中,FAILED：失败,SUCCESS：成功
	AbnormalTableStatus *string `json:"abnormal_table_status,omitempty"`

	// 结果集
	Results *[]interface{} `json:"results,omitempty"`
}

func (o SubInstanceResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubInstanceResult struct{}"
	}

	return strings.Join([]string{"SubInstanceResult", string(data)}, " ")
}
