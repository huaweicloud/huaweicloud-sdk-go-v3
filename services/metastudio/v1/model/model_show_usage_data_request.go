package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUsageDataRequest Request Object
type ShowUsageDataRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 查询时间段开始时间,格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	StartTime string `json:"start_time"`

	// 查询时间段结束时间,格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	EndTime string `json:"end_time"`

	// 用户id(仅开启了子账号隔离功能的主账号携带才生效)
	UserId *string `json:"user_id,omitempty"`

	// 资源类型 * video_time_2d_model：分身数字人视频制作 * video_time_flexus_2d_model：分身数字人视频制作flexus版
	ResourceType *string `json:"resource_type,omitempty"`

	// 业务类型 * LIVE_2D：分身数字人视频直播 * VIDEO_2D：分身数字人视频制作
	BusinessType *string `json:"business_type,omitempty"`

	// 使用量的单位。 * MIN：分钟 * HOUR：小时
	Unit *ShowUsageDataRequestUnit `json:"unit,omitempty"`
}

func (o ShowUsageDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUsageDataRequest struct{}"
	}

	return strings.Join([]string{"ShowUsageDataRequest", string(data)}, " ")
}

type ShowUsageDataRequestUnit struct {
	value string
}

type ShowUsageDataRequestUnitEnum struct {
	MIN  ShowUsageDataRequestUnit
	HOUR ShowUsageDataRequestUnit
}

func GetShowUsageDataRequestUnitEnum() ShowUsageDataRequestUnitEnum {
	return ShowUsageDataRequestUnitEnum{
		MIN: ShowUsageDataRequestUnit{
			value: "MIN",
		},
		HOUR: ShowUsageDataRequestUnit{
			value: "HOUR",
		},
	}
}

func (c ShowUsageDataRequestUnit) Value() string {
	return c.value
}

func (c ShowUsageDataRequestUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUsageDataRequestUnit) UnmarshalJSON(b []byte) error {
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
