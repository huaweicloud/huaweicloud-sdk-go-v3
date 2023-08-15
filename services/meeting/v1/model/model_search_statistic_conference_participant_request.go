package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SearchStatisticConferenceParticipantRequest Request Object
type SearchStatisticConferenceParticipantRequest struct {

	// 查询偏移量。 * 取值：大于等于0，默认值为0 * 大于等于最大条目数量，则返回最后一页数据，页数根据总条目数和limit计算得出
	Offset *int32 `json:"offset,omitempty"`

	// 查询的条目数量。 * 取值：1-500，默认值为20
	Limit *int32 `json:"limit,omitempty"`

	// 查询时间维度，取值： * D: 按日查询 * M: 按月查询
	TimeUnit SearchStatisticConferenceParticipantRequestTimeUnit `json:"timeUnit"`

	// 查询时间范围的开始时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月
	StartTime string `json:"startTime"`

	// 查询时间范围的结束时间，格式根据timeUnit的取值而定。 * timeUnit = D，格式：yyyy-MM-dd，此情况下startTime与endTime间隔最多31日 * timeUnit = M，格式：yyyy-MM，此情况下startTime与endTime间隔最多12个月
	EndTime string `json:"endTime"`

	// 查询分类，取值： * user_participate_info: 用户与会统计数据 * hard_terminal_participate_info: 硬件终端与会统计数据 * participant_type_info: 与会设备统计数据
	Category SearchStatisticConferenceParticipantRequestCategory `json:"category"`
}

func (o SearchStatisticConferenceParticipantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchStatisticConferenceParticipantRequest struct{}"
	}

	return strings.Join([]string{"SearchStatisticConferenceParticipantRequest", string(data)}, " ")
}

type SearchStatisticConferenceParticipantRequestTimeUnit struct {
	value string
}

type SearchStatisticConferenceParticipantRequestTimeUnitEnum struct {
	D SearchStatisticConferenceParticipantRequestTimeUnit
	M SearchStatisticConferenceParticipantRequestTimeUnit
}

func GetSearchStatisticConferenceParticipantRequestTimeUnitEnum() SearchStatisticConferenceParticipantRequestTimeUnitEnum {
	return SearchStatisticConferenceParticipantRequestTimeUnitEnum{
		D: SearchStatisticConferenceParticipantRequestTimeUnit{
			value: "D",
		},
		M: SearchStatisticConferenceParticipantRequestTimeUnit{
			value: "M",
		},
	}
}

func (c SearchStatisticConferenceParticipantRequestTimeUnit) Value() string {
	return c.value
}

func (c SearchStatisticConferenceParticipantRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticConferenceParticipantRequestTimeUnit) UnmarshalJSON(b []byte) error {
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

type SearchStatisticConferenceParticipantRequestCategory struct {
	value string
}

type SearchStatisticConferenceParticipantRequestCategoryEnum struct {
	USER_PARTICIPATE_INFO          SearchStatisticConferenceParticipantRequestCategory
	HARD_TERMINAL_PARTICIPATE_INFO SearchStatisticConferenceParticipantRequestCategory
	PARTICIPANT_TYPE_INFO          SearchStatisticConferenceParticipantRequestCategory
}

func GetSearchStatisticConferenceParticipantRequestCategoryEnum() SearchStatisticConferenceParticipantRequestCategoryEnum {
	return SearchStatisticConferenceParticipantRequestCategoryEnum{
		USER_PARTICIPATE_INFO: SearchStatisticConferenceParticipantRequestCategory{
			value: "user_participate_info",
		},
		HARD_TERMINAL_PARTICIPATE_INFO: SearchStatisticConferenceParticipantRequestCategory{
			value: "hard_terminal_participate_info",
		},
		PARTICIPANT_TYPE_INFO: SearchStatisticConferenceParticipantRequestCategory{
			value: "participant_type_info",
		},
	}
}

func (c SearchStatisticConferenceParticipantRequestCategory) Value() string {
	return c.value
}

func (c SearchStatisticConferenceParticipantRequestCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchStatisticConferenceParticipantRequestCategory) UnmarshalJSON(b []byte) error {
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
