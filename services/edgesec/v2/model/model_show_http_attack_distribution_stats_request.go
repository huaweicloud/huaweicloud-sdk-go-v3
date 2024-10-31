package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHttpAttackDistributionStatsRequest Request Object
type ShowHttpAttackDistributionStatsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 安全统计指标类型。 比如req_num（请求次数），bw（带宽），目前支持req_num
	StatType ShowHttpAttackDistributionStatsRequestStatType `json:"stat_type"`

	// 目前支持action（防护动作）, attack_type（攻击类型）
	GroupBy ShowHttpAttackDistributionStatsRequestGroupBy `json:"group_by"`

	// 起始时间，使用秒级时间戳
	StartTime int64 `json:"start_time"`

	// 结束时间，使用秒级时间戳
	EndTime int64 `json:"end_time"`
}

func (o ShowHttpAttackDistributionStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackDistributionStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackDistributionStatsRequest", string(data)}, " ")
}

type ShowHttpAttackDistributionStatsRequestStatType struct {
	value string
}

type ShowHttpAttackDistributionStatsRequestStatTypeEnum struct {
	REQ_NUM ShowHttpAttackDistributionStatsRequestStatType
	BW      ShowHttpAttackDistributionStatsRequestStatType
}

func GetShowHttpAttackDistributionStatsRequestStatTypeEnum() ShowHttpAttackDistributionStatsRequestStatTypeEnum {
	return ShowHttpAttackDistributionStatsRequestStatTypeEnum{
		REQ_NUM: ShowHttpAttackDistributionStatsRequestStatType{
			value: "req_num",
		},
		BW: ShowHttpAttackDistributionStatsRequestStatType{
			value: "bw",
		},
	}
}

func (c ShowHttpAttackDistributionStatsRequestStatType) Value() string {
	return c.value
}

func (c ShowHttpAttackDistributionStatsRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackDistributionStatsRequestStatType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowHttpAttackDistributionStatsRequestGroupBy struct {
	value string
}

type ShowHttpAttackDistributionStatsRequestGroupByEnum struct {
	ACTION      ShowHttpAttackDistributionStatsRequestGroupBy
	ATTACK_TYPE ShowHttpAttackDistributionStatsRequestGroupBy
}

func GetShowHttpAttackDistributionStatsRequestGroupByEnum() ShowHttpAttackDistributionStatsRequestGroupByEnum {
	return ShowHttpAttackDistributionStatsRequestGroupByEnum{
		ACTION: ShowHttpAttackDistributionStatsRequestGroupBy{
			value: "action",
		},
		ATTACK_TYPE: ShowHttpAttackDistributionStatsRequestGroupBy{
			value: "attack_type",
		},
	}
}

func (c ShowHttpAttackDistributionStatsRequestGroupBy) Value() string {
	return c.value
}

func (c ShowHttpAttackDistributionStatsRequestGroupBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackDistributionStatsRequestGroupBy) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
