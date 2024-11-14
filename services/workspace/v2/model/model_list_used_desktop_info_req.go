package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListUsedDesktopInfoReq 查询使用桌面时长请求。
type ListUsedDesktopInfoReq struct {

	// 桌面id集合。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 开始时间，格式：yyyy-MM-dd（UTC时间，不传查默认最近15天）最多查31天数据。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式：yyyy-MM-dd（UTC时间，不传查默认最近15天）最多查31天数据。
	EndTime *string `json:"end_time,omitempty"`

	// 统计方式，不传则默认按天。可选值为： - DAY: 按天。 - HOUR: 按小时。
	GroupByType *ListUsedDesktopInfoReqGroupByType `json:"group_by_type,omitempty"`

	// 若传桌面的用户名，则查询使用时间只有该用户的使用时间。
	DesktopUsername *string `json:"desktop_username,omitempty"`

	// 从查询结果中的第几条数据开始返回,用于分页查询，取值范围0-2147483647，默认从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 查询结果中想要返回的信息条目数量,用于分页查询，取值范围0-100，默认值100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListUsedDesktopInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsedDesktopInfoReq struct{}"
	}

	return strings.Join([]string{"ListUsedDesktopInfoReq", string(data)}, " ")
}

type ListUsedDesktopInfoReqGroupByType struct {
	value string
}

type ListUsedDesktopInfoReqGroupByTypeEnum struct {
	DAY  ListUsedDesktopInfoReqGroupByType
	HOUR ListUsedDesktopInfoReqGroupByType
}

func GetListUsedDesktopInfoReqGroupByTypeEnum() ListUsedDesktopInfoReqGroupByTypeEnum {
	return ListUsedDesktopInfoReqGroupByTypeEnum{
		DAY: ListUsedDesktopInfoReqGroupByType{
			value: "DAY",
		},
		HOUR: ListUsedDesktopInfoReqGroupByType{
			value: "HOUR",
		},
	}
}

func (c ListUsedDesktopInfoReqGroupByType) Value() string {
	return c.value
}

func (c ListUsedDesktopInfoReqGroupByType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListUsedDesktopInfoReqGroupByType) UnmarshalJSON(b []byte) error {
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
