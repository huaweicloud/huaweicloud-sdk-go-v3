package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSumTableResponse Response Object
type ShowSumTableResponse struct {

	// 结果的ID信息，分页查询的时候带过来。
	ResultId *string `json:"result_id,omitempty"`

	// 数据行列表。
	RowList *[]FrontRow `json:"row_list,omitempty"`

	// 最近一笔数据的时间。
	LatestDataTime *int64 `json:"latest_data_Time,omitempty"`

	// 表格的方向，H：默认，表头横向，V：表头纵向。
	TableDirection *ShowSumTableResponseTableDirection `json:"table_direction,omitempty"`

	// 实际开始的时间。
	RealStartTime *int64 `json:"real_start_time,omitempty"`

	// 实际结束的时间。
	RealEndTime *int64 `json:"real_end_time,omitempty"`

	// 提示信息。
	NoticeMsg *string `json:"notice_msg,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSumTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSumTableResponse struct{}"
	}

	return strings.Join([]string{"ShowSumTableResponse", string(data)}, " ")
}

type ShowSumTableResponseTableDirection struct {
	value string
}

type ShowSumTableResponseTableDirectionEnum struct {
	H ShowSumTableResponseTableDirection
	V ShowSumTableResponseTableDirection
}

func GetShowSumTableResponseTableDirectionEnum() ShowSumTableResponseTableDirectionEnum {
	return ShowSumTableResponseTableDirectionEnum{
		H: ShowSumTableResponseTableDirection{
			value: "H",
		},
		V: ShowSumTableResponseTableDirection{
			value: "V",
		},
	}
}

func (c ShowSumTableResponseTableDirection) Value() string {
	return c.value
}

func (c ShowSumTableResponseTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSumTableResponseTableDirection) UnmarshalJSON(b []byte) error {
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
