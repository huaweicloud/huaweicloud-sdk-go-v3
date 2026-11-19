package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAiOverviewStatisticsRequest Request Object
type ShowAiOverviewStatisticsRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 查询类型（不传时返回全部数据） **约束限制**: 不涉及 **取值范围**: - alarm：告警数量统计 - risk：风险数量统计 - source：智能体防护信息  **默认取值**: 不涉及
	Type *ShowAiOverviewStatisticsRequestType `json:"type,omitempty"`

	// **参数解释**: 开始时间，13位毫秒时间戳 **约束限制**: 需小于等于end_time **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	BeginTime int64 `json:"begin_time"`

	// **参数解释**: 结束时间，13位毫秒时间戳 **约束限制**: 需大于等于begin_time **取值范围**: 最小值0，最大值9223372036854775807 **默认取值**: 不涉及
	EndTime int64 `json:"end_time"`
}

func (o ShowAiOverviewStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiOverviewStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowAiOverviewStatisticsRequest", string(data)}, " ")
}

type ShowAiOverviewStatisticsRequestType struct {
	value string
}

type ShowAiOverviewStatisticsRequestTypeEnum struct {
	ALARM  ShowAiOverviewStatisticsRequestType
	RISK   ShowAiOverviewStatisticsRequestType
	SOURCE ShowAiOverviewStatisticsRequestType
}

func GetShowAiOverviewStatisticsRequestTypeEnum() ShowAiOverviewStatisticsRequestTypeEnum {
	return ShowAiOverviewStatisticsRequestTypeEnum{
		ALARM: ShowAiOverviewStatisticsRequestType{
			value: "alarm",
		},
		RISK: ShowAiOverviewStatisticsRequestType{
			value: "risk",
		},
		SOURCE: ShowAiOverviewStatisticsRequestType{
			value: "source",
		},
	}
}

func (c ShowAiOverviewStatisticsRequestType) Value() string {
	return c.value
}

func (c ShowAiOverviewStatisticsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiOverviewStatisticsRequestType) UnmarshalJSON(b []byte) error {
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
