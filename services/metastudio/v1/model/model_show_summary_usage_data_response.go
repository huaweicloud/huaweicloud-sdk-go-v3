package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSummaryUsageDataResponse Response Object
type ShowSummaryUsageDataResponse struct {

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 业务类型 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作
	BusinessType *string `json:"business_type,omitempty"`

	// 使用个数(视频制作的视频个数,直播的场次)
	Number *int32 `json:"number,omitempty"`

	// 使用量
	Usage *float32 `json:"usage,omitempty"`

	// 使用量的单位。 * MIN：分钟 * HOUR：小时
	Unit *ShowSummaryUsageDataResponseUnit `json:"unit,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSummaryUsageDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSummaryUsageDataResponse struct{}"
	}

	return strings.Join([]string{"ShowSummaryUsageDataResponse", string(data)}, " ")
}

type ShowSummaryUsageDataResponseUnit struct {
	value string
}

type ShowSummaryUsageDataResponseUnitEnum struct {
	MIN  ShowSummaryUsageDataResponseUnit
	HOUR ShowSummaryUsageDataResponseUnit
}

func GetShowSummaryUsageDataResponseUnitEnum() ShowSummaryUsageDataResponseUnitEnum {
	return ShowSummaryUsageDataResponseUnitEnum{
		MIN: ShowSummaryUsageDataResponseUnit{
			value: "MIN",
		},
		HOUR: ShowSummaryUsageDataResponseUnit{
			value: "HOUR",
		},
	}
}

func (c ShowSummaryUsageDataResponseUnit) Value() string {
	return c.value
}

func (c ShowSummaryUsageDataResponseUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSummaryUsageDataResponseUnit) UnmarshalJSON(b []byte) error {
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
