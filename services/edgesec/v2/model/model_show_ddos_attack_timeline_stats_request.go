package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDdosAttackTimelineStatsRequest Request Object
type ShowDdosAttackTimelineStatsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 安全统计指标类型，目前支持bw、pps
	StatType ShowDdosAttackTimelineStatsRequestStatType `json:"stat_type"`

	// bw对应（max_bps、avg_bps）,pps对应（max_pps、avg_pps）
	GroupBy ShowDdosAttackTimelineStatsRequestGroupBy `json:"group_by"`

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`
}

func (o ShowDdosAttackTimelineStatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDdosAttackTimelineStatsRequest struct{}"
	}

	return strings.Join([]string{"ShowDdosAttackTimelineStatsRequest", string(data)}, " ")
}

type ShowDdosAttackTimelineStatsRequestStatType struct {
	value string
}

type ShowDdosAttackTimelineStatsRequestStatTypeEnum struct {
	BW  ShowDdosAttackTimelineStatsRequestStatType
	PPS ShowDdosAttackTimelineStatsRequestStatType
}

func GetShowDdosAttackTimelineStatsRequestStatTypeEnum() ShowDdosAttackTimelineStatsRequestStatTypeEnum {
	return ShowDdosAttackTimelineStatsRequestStatTypeEnum{
		BW: ShowDdosAttackTimelineStatsRequestStatType{
			value: "bw",
		},
		PPS: ShowDdosAttackTimelineStatsRequestStatType{
			value: "pps",
		},
	}
}

func (c ShowDdosAttackTimelineStatsRequestStatType) Value() string {
	return c.value
}

func (c ShowDdosAttackTimelineStatsRequestStatType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDdosAttackTimelineStatsRequestStatType) UnmarshalJSON(b []byte) error {
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

type ShowDdosAttackTimelineStatsRequestGroupBy struct {
	value string
}

type ShowDdosAttackTimelineStatsRequestGroupByEnum struct {
	MAX_BPS ShowDdosAttackTimelineStatsRequestGroupBy
	AVG_BPS ShowDdosAttackTimelineStatsRequestGroupBy
	MAX_PPS ShowDdosAttackTimelineStatsRequestGroupBy
	AVG_PPS ShowDdosAttackTimelineStatsRequestGroupBy
}

func GetShowDdosAttackTimelineStatsRequestGroupByEnum() ShowDdosAttackTimelineStatsRequestGroupByEnum {
	return ShowDdosAttackTimelineStatsRequestGroupByEnum{
		MAX_BPS: ShowDdosAttackTimelineStatsRequestGroupBy{
			value: "max_bps",
		},
		AVG_BPS: ShowDdosAttackTimelineStatsRequestGroupBy{
			value: "avg_bps",
		},
		MAX_PPS: ShowDdosAttackTimelineStatsRequestGroupBy{
			value: "max_pps",
		},
		AVG_PPS: ShowDdosAttackTimelineStatsRequestGroupBy{
			value: "avg_pps",
		},
	}
}

func (c ShowDdosAttackTimelineStatsRequestGroupBy) Value() string {
	return c.value
}

func (c ShowDdosAttackTimelineStatsRequestGroupBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDdosAttackTimelineStatsRequestGroupBy) UnmarshalJSON(b []byte) error {
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
