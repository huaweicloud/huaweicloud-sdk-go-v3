package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListCompetitionWorksRequest struct {
	// 要查询的大赛ID，由大赛平台提供

	CompetitionId int32 `json:"competition_id"`
	// 要查询的大赛阶段ID，由大赛平台提供

	StageId int32 `json:"stage_id"`
	// 查询的截止时间

	ReadTime string `json:"read_time"`
	// 查询的时间范围。day表示以read_time作为结束时间,前一天内作为查询范围,hour表示以read_time作为结束时间,前一小内时作为查询范围。

	TimeUnit *ListCompetitionWorksRequestTimeUnit `json:"time_unit,omitempty"`
	// 作品记录的起始编号,如果不传默认从0开始,offset为0时表示从第一条记录开始

	Offset *int32 `json:"offset,omitempty"`
	// 每页包含的作品记录数,如果不传默认返回100条，并且返回最大条数为100

	Limit *int32 `json:"limit,omitempty"`
	// 需要排序的字段，只支持works_id字段,如果不传则不进行排序

	SortKey *string `json:"sort_key,omitempty"`
	// 排序类型，支持asc|desc，默认为asc升序

	SortDir *ListCompetitionWorksRequestSortDir `json:"sort_dir,omitempty"`
}

func (o ListCompetitionWorksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCompetitionWorksRequest struct{}"
	}

	return strings.Join([]string{"ListCompetitionWorksRequest", string(data)}, " ")
}

type ListCompetitionWorksRequestTimeUnit struct {
	value string
}

type ListCompetitionWorksRequestTimeUnitEnum struct {
	DAY  ListCompetitionWorksRequestTimeUnit
	HOUR ListCompetitionWorksRequestTimeUnit
}

func GetListCompetitionWorksRequestTimeUnitEnum() ListCompetitionWorksRequestTimeUnitEnum {
	return ListCompetitionWorksRequestTimeUnitEnum{
		DAY: ListCompetitionWorksRequestTimeUnit{
			value: "day",
		},
		HOUR: ListCompetitionWorksRequestTimeUnit{
			value: "hour",
		},
	}
}

func (c ListCompetitionWorksRequestTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListCompetitionWorksRequestTimeUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListCompetitionWorksRequestSortDir struct {
	value string
}

type ListCompetitionWorksRequestSortDirEnum struct {
	ASC  ListCompetitionWorksRequestSortDir
	DESC ListCompetitionWorksRequestSortDir
}

func GetListCompetitionWorksRequestSortDirEnum() ListCompetitionWorksRequestSortDirEnum {
	return ListCompetitionWorksRequestSortDirEnum{
		ASC: ListCompetitionWorksRequestSortDir{
			value: "asc",
		},
		DESC: ListCompetitionWorksRequestSortDir{
			value: "desc",
		},
	}
}

func (c ListCompetitionWorksRequestSortDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListCompetitionWorksRequestSortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
