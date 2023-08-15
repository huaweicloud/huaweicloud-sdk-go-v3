package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListIssuesSfV4Request Request Object
type ListIssuesSfV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 偏移量 从0开始
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量 最小1,最大100
	Limit *int32 `json:"limit,omitempty"`

	// 工作项类型
	TrackerId *ListIssuesSfV4RequestTrackerId `json:"tracker_id,omitempty"`

	// 创建工作项的时间(查询的起始时间,查询的结束时间)
	CreatedTimeInterval *string `json:"created_time_interval,omitempty"`

	// 更新工作项的时间(查询的起始时间,查询的结束时间)
	UpdatedTimeInterval *string `json:"updated_time_interval,omitempty"`
}

func (o ListIssuesSfV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssuesSfV4Request struct{}"
	}

	return strings.Join([]string{"ListIssuesSfV4Request", string(data)}, " ")
}

type ListIssuesSfV4RequestTrackerId struct {
	value int32
}

type ListIssuesSfV4RequestTrackerIdEnum struct {
	E_2 ListIssuesSfV4RequestTrackerId
	E_3 ListIssuesSfV4RequestTrackerId
	E_4 ListIssuesSfV4RequestTrackerId
	E_5 ListIssuesSfV4RequestTrackerId
	E_6 ListIssuesSfV4RequestTrackerId
	E_7 ListIssuesSfV4RequestTrackerId
}

func GetListIssuesSfV4RequestTrackerIdEnum() ListIssuesSfV4RequestTrackerIdEnum {
	return ListIssuesSfV4RequestTrackerIdEnum{
		E_2: ListIssuesSfV4RequestTrackerId{
			value: 2,
		}, E_3: ListIssuesSfV4RequestTrackerId{
			value: 3,
		}, E_4: ListIssuesSfV4RequestTrackerId{
			value: 4,
		}, E_5: ListIssuesSfV4RequestTrackerId{
			value: 5,
		}, E_6: ListIssuesSfV4RequestTrackerId{
			value: 6,
		}, E_7: ListIssuesSfV4RequestTrackerId{
			value: 7,
		},
	}
}

func (c ListIssuesSfV4RequestTrackerId) Value() int32 {
	return c.value
}

func (c ListIssuesSfV4RequestTrackerId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListIssuesSfV4RequestTrackerId) UnmarshalJSON(b []byte) error {
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
