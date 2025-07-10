package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDesktopsStatisticsRequest Request Object
type ListDesktopsStatisticsRequest struct {

	// 桌面类型，为空时查所有桌面。 - DEDICATED：普通桌面，包括专享桌面、专属桌面等 - POOLED：池桌面，即桌面池里的桌面
	DesktopType *[]string `json:"desktop_type,omitempty"`

	// 统计类型，为空时仅统计桌面总数 |- - attach-state 按照分配状态统计 - login-state 按照登录状态统计 - run-state 按照运行状态统计。
	StatisticsType *[]ListDesktopsStatisticsRequestStatisticsType `json:"statistics_type,omitempty"`
}

func (o ListDesktopsStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsStatisticsRequest", string(data)}, " ")
}

type ListDesktopsStatisticsRequestStatisticsType struct {
	value string
}

type ListDesktopsStatisticsRequestStatisticsTypeEnum struct {
	ATTACH_STATE ListDesktopsStatisticsRequestStatisticsType
	LOGIN_STATE  ListDesktopsStatisticsRequestStatisticsType
	RUN_STATE    ListDesktopsStatisticsRequestStatisticsType
}

func GetListDesktopsStatisticsRequestStatisticsTypeEnum() ListDesktopsStatisticsRequestStatisticsTypeEnum {
	return ListDesktopsStatisticsRequestStatisticsTypeEnum{
		ATTACH_STATE: ListDesktopsStatisticsRequestStatisticsType{
			value: "attach-state",
		},
		LOGIN_STATE: ListDesktopsStatisticsRequestStatisticsType{
			value: "login-state",
		},
		RUN_STATE: ListDesktopsStatisticsRequestStatisticsType{
			value: "run-state",
		},
	}
}

func (c ListDesktopsStatisticsRequestStatisticsType) Value() string {
	return c.value
}

func (c ListDesktopsStatisticsRequestStatisticsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopsStatisticsRequestStatisticsType) UnmarshalJSON(b []byte) error {
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
