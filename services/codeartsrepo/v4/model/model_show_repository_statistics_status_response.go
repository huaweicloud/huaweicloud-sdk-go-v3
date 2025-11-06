package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepositoryStatisticsStatusResponse Response Object
type ShowRepositoryStatisticsStatusResponse struct {

	// **参数解释：** 是否能被统计。 **取值范围：** - true，可以被统计。 - false，不可被统计。
	CanStatistics *bool `json:"can_statistics,omitempty"`

	// **参数解释：** 是否能被统计。 **取值范围：** - 0，表示用户次数和仓库次数都没用完。 - 1，表示仓库次数用完。 - 2，表示用户次数用完。 - 3, 表示仓库次数和用户次数都用完
	Reason *ShowRepositoryStatisticsStatusResponseReason `json:"reason,omitempty"`

	Event *StatisticEventsDto `json:"event,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryStatisticsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryStatisticsStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryStatisticsStatusResponse", string(data)}, " ")
}

type ShowRepositoryStatisticsStatusResponseReason struct {
	value int32
}

type ShowRepositoryStatisticsStatusResponseReasonEnum struct {
	E_0 ShowRepositoryStatisticsStatusResponseReason
	E_1 ShowRepositoryStatisticsStatusResponseReason
	E_2 ShowRepositoryStatisticsStatusResponseReason
	E_3 ShowRepositoryStatisticsStatusResponseReason
}

func GetShowRepositoryStatisticsStatusResponseReasonEnum() ShowRepositoryStatisticsStatusResponseReasonEnum {
	return ShowRepositoryStatisticsStatusResponseReasonEnum{
		E_0: ShowRepositoryStatisticsStatusResponseReason{
			value: 0,
		}, E_1: ShowRepositoryStatisticsStatusResponseReason{
			value: 1,
		}, E_2: ShowRepositoryStatisticsStatusResponseReason{
			value: 2,
		}, E_3: ShowRepositoryStatisticsStatusResponseReason{
			value: 3,
		},
	}
}

func (c ShowRepositoryStatisticsStatusResponseReason) Value() int32 {
	return c.value
}

func (c ShowRepositoryStatisticsStatusResponseReason) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepositoryStatisticsStatusResponseReason) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
