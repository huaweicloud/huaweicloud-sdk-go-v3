package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowHttpAttackTimelineStatsRequest Request Object
type ShowHttpAttackTimelineStatsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	// 指标类型，当前仅支持req_num
	StatType ShowHttpAttackTimelineStatsRequestStatType `json:"stat_type"`

	// 分组类型，当前仅支持action，attack_category
	GroupBy ShowHttpAttackTimelineStatsRequestGroupBy `json:"group_by"`

	// 分组类型对应的具体的值，不传的话默认总和（例如：action指标类型：log、block、captcha、js_challenge） （例如：attack_category指标类型：cc、access_control、bot、web_app_attack）
	GroupByValue *string `json:"group_by_value,omitempty"`

	// 时间粒度(单位：秒)，不同时间范围有不同的可选时间粒度。[0,1H]，可选时间粒度为1M、5M；(1H,1D]，可选时间粒度为1M、5M、1H；(1D,3D]，可选时间粒度为1M、5M、1H、1D；(3D,7D]，可选时间粒度为5M、1H、1D；(7D,30D]，可选时间粒度为1H、1D。其中M代表分钟，H代表小时，D代表天。
	Interval int32 `json:"interval"`

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`
}

func (o ShowHttpAttackTimelineStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAttackTimelineStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpAttackTimelineStatsRequest", string(data)}, " ")
}

type ShowHttpAttackTimelineStatsRequestStatType struct {
	value string
}

type ShowHttpAttackTimelineStatsRequestStatTypeEnum struct {
	REQ_NUM ShowHttpAttackTimelineStatsRequestStatType
}

func GetShowHttpAttackTimelineStatsRequestStatTypeEnum() ShowHttpAttackTimelineStatsRequestStatTypeEnum {
	return ShowHttpAttackTimelineStatsRequestStatTypeEnum{
		REQ_NUM: ShowHttpAttackTimelineStatsRequestStatType{
			value: "req_num",
		},
	}
}

func (c ShowHttpAttackTimelineStatsRequestStatType) Value() string {
	return c.value
}

func (c ShowHttpAttackTimelineStatsRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackTimelineStatsRequestStatType) UnmarshalJSON(b []byte) error {
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

type ShowHttpAttackTimelineStatsRequestGroupBy struct {
	value string
}

type ShowHttpAttackTimelineStatsRequestGroupByEnum struct {
	ACTION          ShowHttpAttackTimelineStatsRequestGroupBy
	ATTACK_CATEGORY ShowHttpAttackTimelineStatsRequestGroupBy
}

func GetShowHttpAttackTimelineStatsRequestGroupByEnum() ShowHttpAttackTimelineStatsRequestGroupByEnum {
	return ShowHttpAttackTimelineStatsRequestGroupByEnum{
		ACTION: ShowHttpAttackTimelineStatsRequestGroupBy{
			value: "action",
		},
		ATTACK_CATEGORY: ShowHttpAttackTimelineStatsRequestGroupBy{
			value: "attack_category",
		},
	}
}

func (c ShowHttpAttackTimelineStatsRequestGroupBy) Value() string {
	return c.value
}

func (c ShowHttpAttackTimelineStatsRequestGroupBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowHttpAttackTimelineStatsRequestGroupBy) UnmarshalJSON(b []byte) error {
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
