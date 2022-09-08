package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowRawTableResponse struct {
	RowList *[]FrontRow `json:"row_list,omitempty"`

	// 最近一笔数据的时间
	LatestDataTime *string `json:"latest_data_Time,omitempty"`

	// 表格的方向，H：默认，表头横向，V：表头纵向
	TableDirection *ShowRawTableResponseTableDirection `json:"table_direction,omitempty"`

	// 上次请求id
	ResultId *string `json:"result_id,omitempty"`

	// 实际开始的时间，主要用于下一次调用，特别是分页调用的时候传的参数
	RealStartTime *int64 `json:"real_start_time,omitempty"`

	// 实际结束的时间
	RealEndTime    *int64 `json:"real_end_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowRawTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRawTableResponse struct{}"
	}

	return strings.Join([]string{"ShowRawTableResponse", string(data)}, " ")
}

type ShowRawTableResponseTableDirection struct {
	value string
}

type ShowRawTableResponseTableDirectionEnum struct {
	H ShowRawTableResponseTableDirection
	V ShowRawTableResponseTableDirection
}

func GetShowRawTableResponseTableDirectionEnum() ShowRawTableResponseTableDirectionEnum {
	return ShowRawTableResponseTableDirectionEnum{
		H: ShowRawTableResponseTableDirection{
			value: "H",
		},
		V: ShowRawTableResponseTableDirection{
			value: "V",
		},
	}
}

func (c ShowRawTableResponseTableDirection) Value() string {
	return c.value
}

func (c ShowRawTableResponseTableDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRawTableResponseTableDirection) UnmarshalJSON(b []byte) error {
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
